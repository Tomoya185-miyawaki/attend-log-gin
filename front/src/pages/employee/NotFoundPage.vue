<template>
  <HeaderComponent />
  <v-main>
    <v-container>
      <v-row class="text-center">
        <v-col cols="12">
          お探しのページは見つかりませんでした
        </v-col>
        <v-col cols="12">
          <router-link
            elevation="2"
            to="/"
          >
            トップに戻る
          </router-link>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import HeaderComponent from '@/components/layouts/employee/HeaderComponent.vue'

export default defineComponent({
  name: 'EmployeeNotFoundPage',
  components: {
    HeaderComponent,
  },
})
</script>