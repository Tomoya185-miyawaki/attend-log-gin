<template>
  <HeaderComponent />
  <v-main>
    <v-container>
      <h1 class="text-center mb-2">ログイン</h1>
      <v-alert
        dense
        outlined
        type="error"
        class="mb-4"
        v-if="isError"
      >
        認証に失敗しました
      </v-alert>
      <form @submit.prevent="handleSubmit">
        <v-text-field
          v-model="email"
          :error-messages="emailError"
          label="メールアドレス"
          filled
          required
        ></v-text-field>
        <v-text-field
          v-model="password"
          :error-messages="passwordError"
          type="password"
          label="パスワード"
          filled
          required
        ></v-text-field>
        <v-col class="text-right pa-0">
          <v-col class="px-0 pt-0">
            <router-link to="/admin/password-reset">パスワードを忘れた場合はこちら</router-link>
          </v-col>
          <v-btn
            type="submit"
          >
            送信
          </v-btn>
        </v-col>
      </form>
    </v-container>
  </v-main>
  <LoadingComponent :isLoading="isLoading" />
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useField, useForm } from 'vee-validate'
import * as yup from 'yup'
import HeaderComponent from '@/components/layouts/admin/HeaderComponent.vue'
import LoadingComponent from '@/components/parts/LoadingComponent.vue'
import ApiService from '@/services/ApiService'
import router from '@/routes/router'

export default defineComponent({
  name: 'LoginPage',
  components: {
    HeaderComponent,
    LoadingComponent
  },
  setup() {
    let isError = ref<boolean>(false)
    let isLoading = ref<boolean>(false)
    const formSchema = yup.object({
      email: yup.string().email('メールアドレス形式で入力してください').required('メールアドレスは必須項目です'),
      password: yup.string().required('パスワードは必須項目です')
    })
    useForm({ validationSchema: formSchema })
    const { value: email, errorMessage: emailError } = useField<string>('email');
    const { value: password, errorMessage: passwordError } = useField<string>('password');
    const handleSubmit = () => {
      if (
        email.value &&
        !emailError.value?.length &&
        password.value &&
        !passwordError.value?.length
      ) {
        isLoading.value = true
        ApiService.login({
          email: email.value,
          password: password.value
        }).then(() => {
          isLoading.value = false
          router.push({ name: 'employeeList' })
        }).catch(() => {
          isLoading.value = false
          isError.value = true
        })
      }
    }
    return {
      email,
      emailError,
      password,
      passwordError,
      handleSubmit,
      isError,
      isLoading,
    }
  }
})
</script>
