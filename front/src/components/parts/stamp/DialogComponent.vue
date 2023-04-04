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
          :title="dialog === DialogStatus.Create ? '出退勤登録' : '出退勤削除'"
        ></v-toolbar>
        <v-card-text
          v-if="dialog === DialogStatus.Create"
          class="text-center justify-space-around"
        >
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
            v-if="dialog === DialogStatus.Create"
            variant="text"
            color="primary"
            @click="attendRestRegister(status)"
          >
            登録
          </v-btn>
          <v-btn
            v-else-if="dialog === DialogStatus.Delete"
            variant="text"
            color="error"
            @click="attendRestDelete"
          >
            削除
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-col>
</template>

<script lang="ts">
import { defineComponent, watch, ref, toRefs, SetupContext } from 'vue'
import { DialogStatus } from '@/enums/dialog'
import { StampAttendRestStatus } from '@/enums/stamp'
import ApiService from '@/services/ApiService'

type Props = {
  isActive: boolean
  dialogStatus: DialogStatus
  startDate: Date|undefined
  endDate: Date|undefined
  employeeId: string
  stampId: number|undefined
}

export default defineComponent({
  name: 'DialogComponent',
  props: {
    isActive: {
      type: Boolean,
      default: false
    },
    dialogStatus: {
      type: Number,
      required: true,
      validator(value: DialogStatus) {
        return [DialogStatus.Create, DialogStatus.Delete].includes(value)
      }
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
    },
    stampId: {
      type: Number,
      required: false,
      default: undefined
    }
  },
  setup(props: Props, context: SetupContext) {
    const { isActive, dialogStatus } = toRefs(props)
    let isActiveDialog = ref<boolean>(isActive.value)
    let dialog = ref<DialogStatus>(dialogStatus.value)
    let status = ref<StampAttendRestStatus>(StampAttendRestStatus.AttendLeaving)

    watch(isActive, () => {
      isActiveDialog.value = isActive.value
      dialog.value = dialogStatus.value
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
            context.emit('reGetStamp', true)
          })
          .catch(() => {
            alert('出退勤/休憩の登録に失敗しました')
          })
      } else {
        alert('出退勤/休憩の登録に失敗しました')
      }
    }

    const attendRestDelete = async () => {
      isActiveDialog.value = false
      if (props.stampId !== undefined) {
        if (confirm('本当に削除しますか')) {
          await ApiService
            .deleteStamp(props.stampId)
            .then(() => {
              context.emit('reGetStamp', true)
            })
            .catch(() => {
              alert('出退勤/休憩の削除に失敗しました')
            })
        }
      } else {
        alert('出退勤/休憩の削除に失敗しました')
      }
    }

    return {
      isActiveDialog,
      dialog,
      DialogStatus,
      status,
      StampAttendRestStatus,
      attendRestRegister,
      attendRestDelete,
    }
  }
})
</script>