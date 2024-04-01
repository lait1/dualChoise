<script setup lang="ts">
const { t } = useI18n()
import Star from 'assets/svg/Star.svg?component'

const props = defineProps<{
  text: string,
}>()
</script>

<template>
  <header>
    <Star/>
    <h2>{{ props.text }}</h2>
    <Star/>
  </header>
</template>

<style scoped lang="scss">
header{
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 12px;
}
</style>