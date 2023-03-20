<template>
  <HeaderComponent />
  <v-main>
    <v-container>
      <v-col class="pa-0 mb-16">
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
        <p class="text-h4 mb-5">出勤中</p>
        <v-btn
          size="x-large"
          color="success"
          class="w-50"
        >
          出勤する
        </v-btn>
      </v-col>
    </v-container>
  </v-main>
  <LoadingComponent :isLoading="isLoading" />
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { Employee, Stamps } from '@/types/model'
import HeaderComponent from '@/components/layouts/admin/HeaderComponent.vue'
import LoadingComponent from '@/components/parts/LoadingComponent.vue'
import ApiService from '@/services/ApiService'
import router from '@/routes/router'
import { useRoute } from 'vue-router'

export default defineComponent({
  name: 'EmployeeStampCreatePage',
  components: {
    HeaderComponent,
    LoadingComponent,
  },
  setup() {
    let isLoading = ref<boolean>(true)
    let employees = ref<Employee[]>([])
    let stamps = ref<Stamps[]>([])
    const route = useRoute()
    let activeEmployeeId = ref(parseInt(route.params.employeeId as string, 10))

    const getEmployeesData = async () => {
      isLoading.value = true
      await ApiService
        .getEmployees()
        .then(res => {
          employees.value = res.employees
        })
        .getStampsByEmployeeId()
        .then(res => {
          isLoading.value = false
          stamps.value = res.stamps
        })
        .catch(() => {
          isLoading.value = false
        })
    }
    getEmployeesData()

    const moveStampCreatePage = (employeeId: number) => {
      activeEmployeeId.value = employeeId
      router.push({ name: 'stampCreate', params: { employeeId: employeeId }})
    }

    return {
      isLoading,
      employees,
      activeEmployeeId,
      moveStampCreatePage
    }
  }
})
</script>
