<template>
  <v-app-bar app>
    <v-toolbar-title class="pl-0">出退勤管理システム</v-toolbar-title>
  </v-app-bar>
  <LoadingComponent :isLoading="isLoading" />
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import LoadingComponent from '@/components/parts/LoadingComponent.vue'

export default defineComponent({
  name: 'HeaderComponent',
  components: {
    LoadingComponent
  },
  setup() {
    let isLoading = ref<boolean>(false)

    return {
      isLoading
    }
  }
})
</script>