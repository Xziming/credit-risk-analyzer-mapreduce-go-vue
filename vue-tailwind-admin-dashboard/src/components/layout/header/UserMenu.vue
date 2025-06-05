<script setup>
import { ref, onMounted } from 'vue'
import { jwtDecode } from 'jwt-decode'
import { ChevronDownIcon } from 'lucide-vue-next'
import { useRouter } from 'vue-router'

const router = useRouter()
const dropdownOpen = ref(false)
const userName = ref('Admin')

const token = localStorage.getItem('token') || sessionStorage.getItem('token')

onMounted(() => {
  if (token) {
    try {
      const decoded = jwtDecode(token)
      if (decoded?.username) {
        userName.value = decoded.username
      }
    } catch (e) {
      console.error('jwtDecode error:', e)
    }
  } else {
    console.log('No token found in storage')
  }
})

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('username')
  sessionStorage.removeItem('token')
  sessionStorage.removeItem('username')
  router.push('/')
}
</script>

<template>
  <div class="relative">
    <button
      @click.prevent="dropdownOpen = !dropdownOpen"
      class="flex items-center text-gray-700 dark:text-gray-400"
    >
      <span class="mr-3 overflow-hidden rounded-full h-11 w-11">
        <img src="/images/user/owner.jpg" alt="User" />
      </span>
      <span class="block mr-1 font-medium text-theme-sm">{{ userName }}</span>
      <ChevronDownIcon :class="{ 'rotate-180': dropdownOpen }" />
    </button>

    <!-- 下拉菜单 -->
    <div
      v-if="dropdownOpen"
      class="absolute right-0 mt-2 w-48 bg-white border border-gray-200 rounded shadow-lg z-50"
    >
      <div class="px-4 py-2 text-sm text-gray-700">你好，{{ userName }}</div>
      <div class="border-t border-gray-100"></div>
      <a
        href="#"
        @click.prevent="logout"
        class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
      >
        退出登录
      </a>
    </div>
  </div>
</template>

