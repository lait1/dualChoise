<script setup lang="ts">
const { t } = useI18n()
</script>

<template>
  <MainBanner
   :main-text="t('content.main.bannerText')"
   :title="t('content.main.bannerTitle')"
  />
  <PopularQuizzes />
  <PopularCategories />
</template>
