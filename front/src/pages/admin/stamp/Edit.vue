<template>
  <HeaderComponent />
  <v-main>
    <v-container>
      <h2 class="mb-4">{{ employeeName }}の出退勤状況</h2>
      <FullCalendar 
        :options="calendarOptions"
      />
      <DialogComponent
        :isActive="isActive"
        :startDate="startDate"
        :endDate="endDate"
        :employeeId="employeeId"
        @changeActiveDialogFlg="isActive = $event"
        @addAttendRest="reGetStampDetail(employeeId)"
      />
    </v-container>
  </v-main>
  <LoadingComponent :isLoading="isLoading" />
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import '@fullcalendar/core/vdom'
import FullCalendar, { DateSelectArg, EventChangeArg, EventDropArg } from '@fullcalendar/vue3'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import HeaderComponent from '@/components/layouts/admin/HeaderComponent.vue'
import DialogComponent from '@/components/parts/stamp/DialogComponent.vue'
import LoadingComponent from '@/components/parts/LoadingComponent.vue'
import ApiService from '@/services/ApiService'
import { StampAttendRestStatus } from '@/enums/stamp'
import { setEvent, getStatusByTitle } from '@/util/fullCalendar'
import { useRoute } from 'vue-router'
import router from '@/routes/router'

export default defineComponent({
  name: 'StampEditPage',
  components: {
    HeaderComponent,
    LoadingComponent,
    DialogComponent,
    FullCalendar,
  },
  setup() {
    let isLoading = ref<boolean>(true)
    let isActive = ref<boolean>(false)
    let employeeName = ref<string>('')
    let startDate = ref<Date|undefined>(undefined)
    let endDate = ref<Date|undefined>(undefined)
    let calendarOptions = ref<any>(null)
    const route = useRoute()
    const employeeId = route.params.employeeId as string

    calendarOptions.value = {
      plugins: [ timeGridPlugin, interactionPlugin ],
      initialView: 'timeGridWeek',
      headerToolbar: {
        left: 'prev,next',
        center: 'title',
        right: 'timeGridWeek,timeGridDay'
      },
      events: [],
      nowIndicator: true,
      locale: 'ja',
      selectable: true,
      allDaySlot: false,
      select: (selectionInfo: DateSelectArg) => {
        isActive.value = true
        startDate.value = selectionInfo.start
        endDate.value = selectionInfo.end
      },
      eventDrop: (info: EventChangeArg) => {
        const status = getStatusByTitle(info.event.title)
        if (status !== '') {
          updateStamp(info, status)
        }
      },
      eventResize: (info: EventDropArg) => {
        const status = getStatusByTitle(info.event.title)
        if (status !== '') {
          updateStamp(info, status)
        }
      }
    }

    const getStampDetail = async (employeeId: string) => {
      await ApiService
        .getStampDetail(employeeId)
        .then(res => {
          res.stamps.map(stamp => {
            calendarOptions.value.events.push(setEvent(stamp))
          })
        })
        .catch(() => {
          isLoading.value = false
          router.push({ name: 'adminError' })
        })
    }
    getStampDetail(employeeId)

    const reGetStampDetail = (employeeId: string) => {
      calendarOptions.value.events = []
      getStampDetail(employeeId)
    }

    const getEmployeeName = async (employeeId: string) => {
      await ApiService
        .getEmployeesById(employeeId)
        .then(res => {
          employeeName.value = res.employee.name
          isLoading.value = false
        })
        .catch(() => {
          isLoading.value = false
          router.push({ name: 'adminError' })
        })
    }
    getEmployeeName(employeeId)

    const updateStamp = async (info: EventChangeArg | EventDropArg, status: StampAttendRestStatus) => { 
      await ApiService
        .updateStamp({
          id: parseInt(info.event.id, 10),
          status: status,
          stamp_start_date: info.event.start,
          stamp_end_date: info.event.end
        })
        .catch(() => {
          alert('更新に失敗しました。少し待って再度お試しください')
          calendarOptions.value.events = []
          getStampDetail(employeeId)
        })
    }

    return {
      isLoading,
      isActive,
      startDate,
      endDate,
      employeeId,
      employeeName,
      calendarOptions,
      reGetStampDetail,
    }
  }
})
</script>