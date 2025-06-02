<template>
  <div class="grid grid-cols-3 gap-8 w-full">
    <!-- 总客户数 -->
    <div class="rounded-2xl border border-gray-200 bg-white p-5 dark:border-gray-800 dark:bg-white/[0.03] md:p-6">
      <div class="flex items-center justify-center w-12 h-12 bg-gray-100 rounded-xl dark:bg-gray-800">
        <svg class="fill-current text-gray-800 dark:text-white/90" width="24" height="24" viewBox="0 0 24 24" fill="none">
          <path fill-rule="evenodd" clip-rule="evenodd" d="M8.80443 5.60156C7.59109 5.60156 6.60749 6.58517 6.60749 7.79851C6.60749 9.01185 7.59109 9.99545 8.80443 9.99545C10.0178 9.99545 11.0014 9.01185 11.0014 7.79851C11.0014 6.58517 10.0178 5.60156 8.80443 5.60156ZM5.10749 7.79851C5.10749 5.75674 6.76267 4.10156 8.80443 4.10156C10.8462 4.10156 12.5014 5.75674 12.5014 7.79851C12.5014 9.84027 10.8462 11.4955 8.80443 11.4955C6.76267 11.4955 5.10749 9.84027 5.10749 7.79851ZM4.86252 15.3208C4.08769 16.0881 3.70377 17.0608 3.51705 17.8611C3.48384 18.0034 3.5211 18.1175 3.60712 18.2112C3.70161 18.3141 3.86659 18.3987 4.07591 18.3987H13.4249C13.6343 18.3987 13.7992 18.3141 13.8937 18.2112C13.9797 18.1175 14.017 18.0034 13.9838 17.8611C13.7971 17.0608 13.4132 16.0881 12.6383 15.3208C11.8821 14.572 10.6899 13.955 8.75042 13.955C6.81096 13.955 5.61877 14.572 4.86252 15.3208Z" fill="currentColor" />
        </svg>
      </div>
      <div class="flex items-end justify-between mt-5">
        <div>
          <span class="text-sm text-gray-500 dark:text-gray-400">总客户数</span>
          <h4 class="mt-2 font-bold text-gray-800 text-title-sm dark:text-white/90">30000</h4>
        </div>
        <span class="flex items-center gap-1 rounded-full bg-success-50 py-0.5 pl-2 pr-2.5 text-sm font-medium text-success-600 dark:bg-success-500/15 dark:text-success-500">
          0%
        </span>
      </div>
    </div>

    <!-- 轻度逾期客户 -->
    <div class="rounded-2xl border border-gray-200 bg-white p-5 dark:border-gray-800 dark:bg-white/[0.03] md:p-6">
      <div class="flex items-center justify-center w-12 h-12 bg-gray-100 rounded-xl dark:bg-gray-800">
        <svg class="fill-current text-gray-800 dark:text-white/90" width="24" height="24" viewBox="0 0 24 24" fill="none">
          <path fill-rule="evenodd" clip-rule="evenodd"  d="M12 2L22 20H2L12 2ZM12 8V14H14V8H12ZM12 16V18H14V16H12Z" fill="currentColor" />
        </svg>
      </div>
      <div class="flex items-end justify-between mt-5">
        <div>
          <span class="text-sm text-gray-500 dark:text-gray-400">轻度逾期客户</span>
          <h4 class="mt-2 font-bold text-gray-800 text-title-sm dark:text-white/90">6772</h4>
        </div>
        <span class="flex items-center gap-1 rounded-full bg-warning-50 py-0.5 pl-2 pr-2.5 text-sm font-medium text-warning-600 dark:bg-warning-500/15 dark:text-warning-500">
          22.6%
        </span>
      </div>
    </div>

    <!-- 高风险客户 -->
    <div class="rounded-2xl border border-gray-200 bg-white p-5 dark:border-gray-800 dark:bg-white/[0.03] md:p-6">
      <div class="flex items-center justify-center w-12 h-12 bg-gray-100 rounded-xl dark:bg-gray-800">
        <svg class="fill-current text-gray-800 dark:text-white/90" width="24" height="24" viewBox="0 0 24 24" fill="none">
          <path fill-rule="evenodd" clip-rule="evenodd" d="M12 2C10.3431 2 9 3.34315 9 5V5.58579C9 6.10948 8.78929 6.60228 8.41421 6.97631L6.70711 8.68342C6.25971 9.13082 6 9.74264 6 10.3787V16L4 18V19H20V18L18 16V10.3787C18 9.74264 17.7403 9.13082 17.2929 8.68342L15.5858 6.97631C15.2107 6.60228 15 6.10948 15 5.58579V5C15 3.34315 13.6569 2 12 2ZM12 22C13.1046 22 14 21.1046 14 20H10C10 21.1046 10.8954 22 12 22Z" fill="currentColor" />
        </svg>
      </div>
      <div class="flex items-end justify-between mt-5">
        <div>
          <span class="text-sm text-gray-500 dark:text-gray-400">高风险客户</span>
          <h4 class="mt-2 font-bold text-gray-800 text-title-sm dark:text-white/90">3297</h4>
        </div>
        <span class="flex items-center gap-1 rounded-full bg-error-50 py-0.5 pl-2 pr-2.5 text-sm font-medium text-error-600 dark:bg-error-500/15 dark:text-error-500">
          11.0%
        </span>
      </div>
    </div>
  </div>
</template>

