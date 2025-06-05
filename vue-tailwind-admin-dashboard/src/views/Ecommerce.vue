<template>
<admin-layout>
<div class="grid grid-cols-12 gap-4 md:gap-6">
<div class="col-span-12">
<ecommerce-metrics />
</div>
<div class="col-span-12 xl:col-span-4">
<monthly-Target />
</div>

<div class="col-span-12 xl:col-span-8">
<statistics-chart />
</div>

<div class="col-span-12 xl:col-span-5">
<customer-demographic />
</div>

<div class="col-span-12 xl:col-span-7">
<recent-orders />
</div>
</div>
</admin-layout>
</template>

<script>
import { jwtDecode } from 'jwt-decode'
import { useRouter } from 'vue-router'
import { onMounted } from 'vue'


import AdminLayout from '../components/layout/AdminLayout.vue'
import EcommerceMetrics from '../components/ecommerce/EcommerceMetrics.vue'
import MonthlyTarget from '../components/ecommerce/MonthlyTarget.vue'
import StatisticsChart from '../components/ecommerce/StatisticsChart.vue'
export default {
components: {
                AdminLayout,
                    EcommerceMetrics,
                    MonthlyTarget,
                    StatisticsChart,
            },
name: 'Ecommerce',
      setup() {
          const router = useRouter()

              // 安全校验（防止 token 丢失时异常渲染）
              const token = localStorage.getItem('token') || sessionStorage.getItem('token')
              if (!token) {
                  router.push('/')
              } else {
                  try {
                      jwtDecode(token) // 可加 decode 验证
                  } catch (err) {
                      console.error('无效 token:', err)
                          router.push('/')
                  }
              }
      },
}
onMounted(() => {
        console.log('当前 token:', localStorage.getItem('token') || sessionStorage.getItem('token'))
        })

</script>
