<template>
<div class="flex h-screen bg-gray-100">
<!-- 侧边栏 -->
<Sidebar />

<!-- 主内容 -->
<div class="flex-1 flex flex-col overflow-hidden">
<Header />

<!-- 内容区域 -->
<main class="flex-1 overflow-y-auto p-6 space-y-6">
<div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 gap-6">
<GenderChart />
<MarriageBarChart />
<EducationBarChart />
<LimitBalChart />
</div>
</main>
</div>
</div>
</template>

<script setup>
import Sidebar from '../layout/Sidebar.vue'
import Header from '../layout/Header.vue'
import GenderChart from '../components/GenderChart.vue'
import MarriageBarChart from '../components/MarriageBarChart.vue'
import EducationBarChart from '../components/EducationBarChart.vue'
import LimitBalChart from '../components/LimitBalChart.vue'
import AgeDefaultsChart from '../components/AgeDefaultsChart.vue'
import GenderDefaultsChart from '../components/GenderDefaultsChart.vue'
import EducationDefaultsChart from '../components/EducationDefaultsChart.vue'
import MarriageDefaultsChart from '../components/MarriageDefaultsChart.vue'
import LimitBalDefaultsChart from '../components/LimitBalDefaultsChart.vue'
</script>

