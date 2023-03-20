<template>
  <HeaderComponent />
  <v-main>
    <v-container>
      <v-col class="pa-0">
        <v-btn 
          v-for="employee in employees"
          :key="employee.id"
          variant="outlined"
          class="mb-2 mr-2"
          @click="moveStampCreatePage(employee.id)"
        >
          社員番号：{{ employee.id }} {{ employee.name }}
        </v-btn>
      </v-col>
    </v-container>
  </v-main>
  <LoadingComponent :isLoading="isLoading" />
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { Employee } from '@/types/model'
import HeaderComponent from '@/components/layouts/admin/HeaderComponent.vue'
import LoadingComponent from '@/components/parts/LoadingComponent.vue'
import ApiService from '@/services/ApiService'
import router from '@/routes/router'

export default defineComponent({
  name: 'EmployeeStampListPage',
  components: {
    HeaderComponent,
    LoadingComponent,
  },
  setup() {
    let isLoading = ref<boolean>(true)
    let employees = ref<Employee[]>([])

    const getEmployees = async () => {
      isLoading.value = true
      await ApiService
        .getEmployees()
        .then(res => {
          isLoading.value = false
          employees.value = res.employees
        })
        .catch(() => {
          isLoading.value = false
        })
    }
    getEmployees()

    const moveStampCreatePage = (employeeId: number) => {
      router.push({ name: 'stampCreate', params: { employeeId: employeeId }})
    }

    return {
      isLoading,
      employees,
      moveStampCreatePage
    }
  }
})
</script>
