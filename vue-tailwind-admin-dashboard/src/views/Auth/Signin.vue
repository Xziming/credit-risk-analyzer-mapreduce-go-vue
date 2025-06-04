<template>
<div class="min-h-screen flex items-center justify-center bg-gray-100">
<div class="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
<h2 class="text-2xl font-bold mb-6 text-center">登录</h2>
<form @submit.prevent="handleLogin">
<div class="mb-4">
<label class="block mb-1 text-gray-700">用户名</label>
<input
v-model="form.username"
type="text"
class="w-full border rounded px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
required
/>
</div>
<div class="mb-6">
<label class="block mb-1 text-gray-700">密码</label>
<input
v-model="form.password"
type="password"
class="w-full border rounded px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
required
/>
</div>
<button
type="submit"
class="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600"
>
登录
</button>
<p v-if="error" class="mt-4 text-center text-red-500">{{ error }}</p>
</form>

<div class="mb-6" text-center>
<p class="mt-6 text-center text-gray-600">
还没有账号？
<button @click="goToRegister" class="text-blue-500 hover:underline"></button>
</p>
</div>
</div>
</div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()
    const form = ref({ username: '', password: '' })
    const error = ref('')

    const handleLogin = async () => {
        try {
            const res = await axios.post('http://zzh.hengtai119.cn/api/login', form.value)
                const { token, user } = res.data

                localStorage.setItem('token', token)
                localStorage.setItem('user', JSON.stringify(user))

                router.push('/') // 或你项目的首页路径
        } catch (err) {
            error.value = err.response?.data?.message || '用户名或密码错误'
        }
    }
const goToRegister = () => {
    router.push('/signup')
}
</script>

