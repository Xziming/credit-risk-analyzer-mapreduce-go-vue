<template>
<AdminLayout>
<PageBreadcrumb :pageTitle="currentPageTitle" />
<div class="space-y-5 sm:space-y-6">
<ComponentCard title="Over Due Chart"> <OverDueChart /> </ComponentCard>
</div>
</AdminLayout>
</template>

<script setup>
import OverDueChart from "@/components/charts/MoreChart/OverDueChart.vue";
import ComponentCard from "@/components/common/ComponentCard.vue";
import PageBreadcrumb from "@/components/common/PageBreadcrumb.vue";
import AdminLayout from "@/components/layout/AdminLayout.vue";
import { ref } from "vue";
const currentPageTitle = ref("Over Due Chart");
</script>
