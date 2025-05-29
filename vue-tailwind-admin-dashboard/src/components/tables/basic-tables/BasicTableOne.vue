<template>
<div class="p-4 bg-white dark:bg-gray-900 rounded-xl shadow-md">
<h2 class="text-lg font-semibold text-gray-800 dark:text-white mb-4">客户信息表</h2>
<div class="overflow-x-auto custom-scrollbar">
<table class="min-w-full text-sm text-left">
<thead class="bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-300">
<tr>
<th v-for="header in headers" :key="header" class="px-4 py-2 whitespace-nowrap border-b border-gray-300 dark:border-gray-700">
{{ header }}
</th>
</tr>
</thead>
<tbody class="divide-y divide-gray-200 dark:divide-gray-700">
<tr v-for="client in paginatedData" :key="client.ID" class="hover:bg-gray-50 dark:hover:bg-gray-800">
<td v-for="key in keys" :key="key" class="px-4 py-2 whitespace-nowrap text-gray-700 dark:text-gray-200">
{{ client[key] }}
</td>
</tr>
</tbody>
</table>
</div>

<!-- 分页 -->
<div class="flex justify-end items-center mt-4 space-x-2 text-sm text-gray-600 dark:text-gray-300">
<button @click="prevPage" :disabled="currentPage === 1" class="px-3 py-1 rounded border dark:border-gray-600 hover:bg-gray-100 dark:hover:bg-gray-700">
上一页
</button>
<span>第 {{ currentPage }} 页 / 共 {{ totalPages }} 页</span>
<button @click="nextPage" :disabled="currentPage === totalPages" class="px-3 py-1 rounded border dark:border-gray-600 hover:bg-gray-100 dark:hover:bg-gray-700">
下一页
</button>
</div>
</div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import axios from 'axios'

// 所有字段的顺序
const keys = [
    'client_id', 'limit_bal', 'sex', 'education', 'marriage', 'age',
    'pay_0', 'pay_2', 'pay_3', 'pay_4', 'pay_5', 'pay_6',
    'bill_amt1', 'bill_amt2', 'bill_amt3', 'bill_amt4', 'bill_amt5', 'bill_amt6',
    'pay_amt1', 'pay_amt2', 'pay_amt3', 'pay_amt4', 'pay_amt5', 'pay_amt6',
    'defaulted'
]

// 表头显示名（可根据需求修改）
const headers = [
    'ID', '额度', '性别', '教育', '婚姻', '年龄',
    '9月还款状态', '8月还款状态', '7月还款状态', '6月还款状态', '5月还款状态', '4月还款状态',
    '9月账单', '8月账单', '7月账单', '6月账单', '5月账单', '4月账单',
    '9月还款额', '8月还款额', '7月还款额', '6月还款额', '5月还款额', '4月还款额',
    '是否违约'
]

    const clients = ref([])
const currentPage = ref(1)
    const pageSize = 10

    const paginatedData = computed(() => {
            const start = (currentPage.value - 1) * pageSize
            return clients.value.slice(start, start + pageSize)
            })

const totalPages = computed(() => Math.ceil(clients.value.length / pageSize))

    const fetchData = async () => {
        try {
            const response = await axios.get('http://zzh.hengtai119.cn/api/clientsinfo')
                clients.value = response.data.slice(0, 30000)
        } catch (error) {
            console.error('获取客户信息失败:', error)
        }
    }

const prevPage = () => {
    if (currentPage.value > 1) currentPage.value--
}

const nextPage = () => {
    if (currentPage.value < totalPages.value) currentPage.value++
}

onMounted(fetchData)
    </script>

    <style scoped>
    .custom-scrollbar::-webkit-scrollbar {
height: 6px;
    }
.custom-scrollbar::-webkit-scrollbar-thumb {
    background-color: rgba(100, 100, 100, 0.3);
    border-radius: 4px;
}
</style>

