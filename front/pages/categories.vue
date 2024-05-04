<script setup lang="ts">
import type {ICategoryInfo} from "#build/types/category.type";
import type {IQuiz} from "#build/types/quiz.type";

const { t } = useI18n()

useHead({
    title: 'YOUR Best categories'
})
const categories = ref<ICategoryInfo[]>([])
const { data } = await useServerFetch<IQuiz[]>("/api/get-categories")

categories.value = data.value
</script>
<template>
    <section class="categories">
        <div class="wrap">
            <div class="categories__title">
                <h1>{{ t('content.main.categories') }}</h1>
            </div>
            <Categories :categories="categories" />
        </div>
    </section>
</template>

<style scoped lang="scss">
.categories__title{
    text-align: center;
    margin-bottom: 52px;
}
</style>