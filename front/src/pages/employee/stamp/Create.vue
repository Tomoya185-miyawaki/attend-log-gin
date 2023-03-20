<template>
  <HeaderComponent />
  <v-main>
    <v-container>
      <v-col class="pa-0 mb-16">
        <v-alert
          v-if="createMessage !== ''"
          class="mb-4"
          :type="isSuccessCreate ? 'success' : 'error'"
          :title="createMessage"
        ></v-alert>
        <v-btn 
          v-for="employee in employees"
          :key="employee.id"
          variant="outlined"
          class="mb-2 mr-2"
          :color="activeEmployeeId === employee.id ? 'info' : ''"
          @click="moveStampCreatePage(employee.id)"
        >
          社員番号：{{ employee.id }} {{ employee.name }}
        </v-btn>
      </v-col>
      <v-col class="pa-0 text-center">
        <p class="text-h4 mb-5">
          {{ attendRestStatusText }}
        </p>
        <v-btn
          v-if="attendRestStatusButtonText !== ''"
          size="x-large"
          color="success"
          class="w-50"
          @click="createStamp"
        >
          {{ attendRestStatusButtonText }}する
        </v-btn>
      </v-col>
    </v-container>
  </v-main>
  <LoadingComponent :isLoading="isLoading" />
</template>

<script lang="ts">
import { defineComponent, ref, watch } from 'vue'
import { StampStatus } from '@/enums/stamp'
import { Employee } from '@/types/model'
import { getAttendRestStatus } from '@/util/stamp'
import HeaderComponent from '@/components/layouts/admin/HeaderComponent.vue'
import LoadingComponent from '@/components/parts/LoadingComponent.vue'
import ApiService from '@/services/ApiService'
import router from '@/routes/router'
import { useRoute } from 'vue-router'
import { getFormatToday } from '@/util/date'

export default defineComponent({
  name: 'EmployeeStampCreatePage',
  components: {
    HeaderComponent,
    LoadingComponent,
  },
  setup() {
    let isLoading = ref<boolean>(true)
    let isSuccessCreate = ref<boolean>(false)
    let createMessage = ref<string>('')
    let employees = ref<Employee[]>([])
    let attendRestStatusText = ref<string|StampStatus>('')
    let attendRestStatusButtonText = ref<string|StampStatus>('')
    let attendRestStatus = ref<string|StampStatus>('')
    const route = useRoute()
    const employeeId = route.params.employeeId as string
    let activeEmployeeId = ref(parseInt(employeeId, 10))
    const today = getFormatToday('-', '-')

    const getEmployees = async () => {
      isLoading.value = true
      await ApiService
        .getEmployees()
        .then(res => {
          employees.value = res.employees
        })
    }
    getEmployees()

    const getStampsByEmployeeId = async () => {
      isLoading.value = true
      await ApiService
        .getStampsByEmployeeId(activeEmployeeId.value, today)
        .then(res => {
          isLoading.value = false
          const attendRestText = getAttendRestStatus(res.stamps)
          attendRestStatusText.value = attendRestText[0]
          attendRestStatusButtonText.value = attendRestText[1]
          attendRestStatus.value = attendRestText[2]
        })
        .catch(() => {
          isLoading.value = false
        })
    }
    getStampsByEmployeeId()

    const moveStampCreatePage = (employeeId: number) => {
      activeEmployeeId.value = employeeId
      router.push({ name: 'stampCreate', params: { employeeId: employeeId }})
    }

    const createStamp = async () => {
      isLoading.value = true
      await ApiService
        .createStamp({
          employeeId: activeEmployeeId.value,
          status: attendRestStatus.value as StampStatus,
          today: today
        })
        .then(() => {
          isLoading.value = false
          isSuccessCreate.value = true
          createMessage.value = `${attendRestStatusButtonText.value}の登録に成功しました`
          getStampsByEmployeeId()
        })
        .catch(() => {
          isLoading.value = false
          isSuccessCreate.value = false
          createMessage.value = `${attendRestStatusButtonText.value}の登録に失敗しました`
        })
    }

    watch(activeEmployeeId, () => {
      isSuccessCreate.value = false
      createMessage.value = ''
      getStampsByEmployeeId()
    })

    return {
      isLoading,
      employees,
      attendRestStatusText,
      attendRestStatusButtonText,
      activeEmployeeId,
      moveStampCreatePage,
      createStamp,
      isSuccessCreate,
      createMessage
    }
  }
})
</script>
