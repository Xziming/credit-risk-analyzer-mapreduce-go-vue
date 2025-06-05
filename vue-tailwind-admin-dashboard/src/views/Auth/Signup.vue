<template>
<FullScreenLayout>
<div class="relative p-6 bg-white z-1 dark:bg-gray-900 sm:p-0">
<div class="relative flex flex-col justify-center w-full h-screen lg:flex-row dark:bg-gray-900">
<!-- 左侧注册表单 -->
<div class="flex flex-col flex-1 w-full lg:w-1/2">
<div class="w-full max-w-md pt-10 mx-auto">
<router-link
to="/"
class="inline-flex items-center text-sm text-gray-500 transition-colors hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-300"
>
<svg class="stroke-current" xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none">
<path d="M12.7083 5L7.5 10.2083L12.7083 15.4167" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
</svg>
Back to dashboard
</router-link>
</div>

<div class="flex flex-col justify-center flex-1 w-full max-w-md mx-auto">
<div class="mb-5 sm:mb-8">
<h1 class="mb-2 font-semibold text-gray-800 text-title-sm dark:text-white/90 sm:text-title-md">
Sign Up
</h1>
<p class="text-sm text-gray-500 dark:text-gray-400">
Enter your username and password to sign up!
</p>
</div>

<form @submit.prevent="handleSubmit">
<div class="space-y-5">
<!-- Username -->
<div>
<label for="username" class="mb-1.5 block text-sm font-medium text-gray-700 dark:text-gray-400">
Username<span class="text-error-500">*</span>
</label>
<input
v-model="username"
type="text"
id="username"
placeholder="Enter your username"
class="dark:bg-dark-900 h-11 w-full rounded-lg border border-gray-300 bg-transparent px-4 py-2.5 text-sm text-gray-800 shadow-theme-xs placeholder:text-gray-400 focus:border-brand-300 focus:outline-hidden focus:ring-3 focus:ring-brand-500/10 dark:border-gray-700 dark:bg-gray-900 dark:text-white/90 dark:placeholder:text-white/30 dark:focus:border-brand-800"
/>
</div>

<!-- Password -->
<div>
<label for="password" class="mb-1.5 block text-sm font-medium text-gray-700 dark:text-gray-400">
Password<span class="text-error-500">*</span>
</label>
<div class="relative">
<input
v-model="password"
:type="showPassword ? 'text' : 'password'"
id="password"
placeholder="Enter your password"
class="dark:bg-dark-900 h-11 w-full rounded-lg border border-gray-300 bg-transparent py-2.5 pl-4 pr-11 text-sm text-gray-800 shadow-theme-xs placeholder:text-gray-400 focus:border-brand-300 focus:outline-hidden focus:ring-3 focus:ring-brand-500/10 dark:border-gray-700 dark:bg-gray-900 dark:text-white/90 dark:placeholder:text-white/30 dark:focus:border-brand-800"
/>
<span
@click="togglePasswordVisibility"
class="absolute z-30 text-gray-500 -translate-y-1/2 cursor-pointer right-4 top-1/2 dark:text-gray-400"
>
<svg v-if="!showPassword" class="fill-current" width="20" height="20" viewBox="0 0 20 20" fill="none">
<path fill-rule="evenodd" clip-rule="evenodd"
d="M10.0002 13.8619C7.23361 13.8619 4.86803 12.1372 3.92328 9.70241C4.86804 7.26761 7.23361 5.54297 10.0002 5.54297C12.7667 5.54297 15.1323 7.26762 16.0771 9.70243C15.1323 12.1372 12.7667 13.8619 10.0002 13.8619ZM10.0002 4.04297C6.48191 4.04297 3.49489 6.30917 2.4155 9.4593C2.3615 9.61687 2.3615 9.78794 2.41549 9.94552C3.49488 13.0957 6.48191 15.3619 10.0002 15.3619C13.5184 15.3619 16.5055 13.0957 17.5849 9.94555C17.6389 9.78797 17.6389 9.6169 17.5849 9.45932C16.5055 6.30919 13.5184 4.04297 10.0002 4.04297Z"
fill="#98A2B3"/>
</svg>
<svg v-else class="fill-current" width="20" height="20" viewBox="0 0 20 20" fill="none">
<path fill-rule="evenodd" clip-rule="evenodd"
d="M4.63803 3.57709C4.34513 3.2842 3.87026 3.2842 3.57737 3.57709C3.28447 3.86999 3.28447 4.34486 3.57737 4.63775L4.85323 5.91362C3.74609 6.84199 2.89363 8.06395 2.4155 9.45936C2.3615 9.61694 2.3615 9.78801 2.41549 9.94558C3.49488 13.0957 6.48191 15.3619 10.0002 15.3619C11.255 15.3619 12.4422 15.0737 13.4994 14.5598L15.3625 16.4229C15.6554 16.7158 16.1302 16.7158 16.4231 16.4229C16.716 16.13 16.716 15.6551 16.4231 15.3622L4.63803 3.57709Z"
fill="#98A2B3"/>
</svg>
</span>
</div>
</div>
<!-- Confirm Password -->
<div>
<label
for="confirmPassword"
class="mb-1.5 block text-sm font-medium text-gray-700 dark:text-gray-400"
>
Confirm Password<span class="text-error-500">*</span>
</label>
<input
v-model="confirmPassword"
:type="showPassword ? 'text' : 'password'"
id="confirmPassword"
placeholder="Confirm your password"
class="dark:bg-dark-900 h-11 w-full rounded-lg border border-gray-300 bg-transparent px-4 py-2.5 text-sm text-gray-800 shadow-theme-xs placeholder:text-gray-400 focus:border-brand-300 focus:outline-hidden focus:ring-3 focus:ring-brand-500/10 dark:border-gray-700 dark:bg-gray-900 dark:text-white/90 dark:placeholder:text-white/30 dark:focus:border-brand-800"
/>
</div>

<!-- Checkbox -->
<div>
<label for="agree" class="flex items-start text-sm font-normal text-gray-700 cursor-pointer dark:text-gray-400">
<input v-model="agreeToTerms" type="checkbox" id="agree" class="mr-2 mt-1"/>
By creating an account you agree to our <span class="text-gray-800 dark:text-white ml-1">Terms and Privacy Policy</span>
</label>
</div>

<!-- Submit Button -->
<div>
<button
type="submit"
class="flex items-center justify-center w-full px-4 py-3 text-sm font-medium text-white transition rounded-lg bg-brand-500 shadow-theme-xs hover:bg-brand-600"
>
Sign Up
</button>
</div>
</div>
</form>

<div class="mt-5 text-sm text-center text-gray-700 dark:text-gray-400">
Already have an account?
<router-link to="/" class="text-brand-500 hover:text-brand-600 dark:text-brand-400">Sign In</router-link>
</div>
</div>
</div>

<!-- 右侧图形背景 -->
<div class="relative items-center hidden w-full h-full lg:w-1/2 bg-brand-950 dark:bg-white/5 lg:grid">
<div class="flex items-center justify-center z-1">
<common-grid-shape />
<div class="flex flex-col items-center max-w-xs">
<router-link to="/" class="block mb-4">
<img width="231" height="48" src="/images/logo/auth-logo.svg" alt="Logo" />
</router-link>
<p class="text-center text-gray-400 dark:text-white/60">
Free and Open-Source Tailwind CSS Admin Dashboard Template
</p>
</div>
</div>
</div>
</div>
</div>
</FullScreenLayout>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import axios, { AxiosError } from 'axios'
import FullScreenLayout from '@/components/layout/FullScreenLayout.vue'
import CommonGridShape from '@/components/common/CommonGridShape.vue'

const username = ref('')
const password = ref('')
    const confirmPassword = ref('')
    const showPassword = ref(false)
const agreeToTerms = ref(false)

    const togglePasswordVisibility = () => {
        showPassword.value = !showPassword.value
    }

const handleSubmit = async () => {
    if (!username.value || !password.value) {
        alert('Please enter both username and password.')
            return
    }

    if (password.value !== confirmPassword.value) {
        alert('Passwords do not match.')
            return
    }

    if (!agreeToTerms.value) {
        alert('Please agree to the terms.')
            return
    }

    try {
        const response = await axios.post('http://zzh.hengtai119.cn/api/register', {
username: username.value,
password: password.value,
})

if (response.status === 200) {
    alert('Registration successful!')
        // router.push('/')
} else {
    alert('Registration failed.')
}
} catch (err: unknown) {
    const error = err as AxiosError<any>
        alert('Error: ' + (error.response?.data?.message || error.message))
}
}

</script>

