<template>
  <v-col cols="auto">
    <v-dialog
      transition="dialog-top-transition"
      width="auto"
      v-model="isActiveDialog"
      class="h-50 w-50"
    >
      <v-card width="300">
        <v-toolbar
          color="primary"
          title="出退勤登録"
        ></v-toolbar>
        <v-card-text class="text-center justify-space-around">
          <v-btn
            variant="flat"
            color="secondary"
            class="mr-4"
            @click="status = StampAttendRestStatus.AttendLeaving"
          >
            勤務
          </v-btn>
          <v-btn
            variant="flat"
            color="info"
            @click="status = StampAttendRestStatus.Rest"
          >
            休憩
          </v-btn>
        </v-card-text>
        <v-card-actions class="justify-space-around">
          <v-btn
            variant="text"
            @click="isActiveDialog = false"
          >
            閉じる
          </v-btn>
          <v-btn
            variant="text"
            color="primary"
            @click="attendRestRegister(status)"
          >
            登録
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-col>
</template>

<script lang="ts">
import { defineComponent, watch, ref, toRefs, SetupContext } from 'vue'
import { StampAttendRestStatus } from '@/enums/stamp'
import ApiService from '@/services/ApiService'

type Props = {
  isActive: boolean
  startDate: Date|undefined
  endDate: Date|undefined
  employeeId: string
}

export default defineComponent({
  name: 'DialogComponent',
  props: {
    isActive: {
      type: Boolean,
      default: false
    },
    startDate: {
      type: Date,
      required: false,
      default: undefined
    },
    endDate: {
      type: Date,
      required: false,
      default: undefined
    },
    employeeId: {
      type: String,
      required: true
    }
  },
  setup(props: Props, context: SetupContext) {
    const { isActive } = toRefs(props)
    let isActiveDialog = ref<boolean>(isActive.value)
    let status = ref<StampAttendRestStatus>(StampAttendRestStatus.AttendLeaving)

    watch(isActive, () => {
      isActiveDialog.value = isActive.value
    })
    watch(isActiveDialog, () => {
      context.emit('changeActiveDialogFlg', isActiveDialog.value)
    })

    const attendRestRegister = async (status: StampAttendRestStatus) => {
      isActiveDialog.value = false

      if (props.startDate !== undefined || props.endDate !== undefined) {
        await ApiService
          .adminCreateStamp({
            employeeId: parseInt(props.employeeId, 10),
            status: status,
            stamp_start_date: props.startDate as Date,
            stamp_end_date: props.endDate as Date
          })
          .then(() => {
            context.emit('addAttendRest', true)
          })
          .catch(() => {
            alert('出退勤/休憩の登録に失敗しました')
          })
      } else {
        alert('出退勤/休憩の登録に失敗しました')
      }
    }
    return {
      isActiveDialog,
      status,
      StampAttendRestStatus,
      attendRestRegister,
    }
  }
})
</script>