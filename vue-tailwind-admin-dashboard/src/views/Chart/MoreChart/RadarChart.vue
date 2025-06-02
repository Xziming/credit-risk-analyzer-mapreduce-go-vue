<template>
<AdminLayout>
<PageBreadcrumb :pageTitle="currentPageTitle" />
<div class="space-y-5 sm:space-y-6">
<ComponentCard title="Radar Chart"> <RadarChart /> </ComponentCard>
</div>
</AdminLayout>
</template>

<script setup>
import RadarChart from "@/components/charts/MoreChart/RadarChart.vue";
import ComponentCard from "@/components/common/ComponentCard.vue";
import PageBreadcrumb from "@/components/common/PageBreadcrumb.vue";
import AdminLayout from "@/components/layout/AdminLayout.vue";
import { ref } from "vue";
const currentPageTitle = ref("Radar Chart");
</script>
