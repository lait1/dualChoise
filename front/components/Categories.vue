<script setup lang="ts">
import type {ICategoryInfo} from "#build/types/category.type";
const { t } = useI18n()
const props = withDefaults(defineProps<{
    categories: ICategoryInfo[]
}>(), {
    categories: []
})

</script>

<template>
    <div class="category-container">
        <CategoryItem
         v-for="category in props.categories"
         :key="category.id"
         :category="category"
         @click="navigateTo(
          `/category/${category.id}`,{ external: true }
        )"
        />
    </div>
</template>

<style scoped lang="scss">
.category-container{
    display: flex;
    justify-content: space-between;
    gap: 16px;
    flex-wrap: wrap;
}
</style>