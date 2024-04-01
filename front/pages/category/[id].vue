<script setup lang="ts">
import type { IQuiz } from "#build/types/quiz.type";
import type {ICategory} from "#build/types/category.type";
const { t } = useI18n()
const { params } = useRoute()
const { isDesktop } = useDevice();
useHead({
    title: `Category id = ${params.id}`
})
const imageSize = isDesktop ? 200 : 250
const info = ref<ICategory>(null)
const { data } = await useAPIFetch<IQuiz[]>(`/api/get-quizzes-by-category/${params.id}`)

info.value = data.value
</script>
<template>
    <section class="categories">
        <div class="wrap">
            <h1 class="categories__title">{{ t(info.categoryName) }}</h1>
            <div class="categories__quizzes">
                <Quiz
                 class="categories__quiz"
                 v-for="quiz in info.quizzes"
                 :key="quiz.id"
                 :quiz="quiz"
                 :size="imageSize"
                 @click="navigateTo(
                  `/start-quiz/${quiz.id}`
                    )"
                />
            </div>
        </div>
    </section>
</template>

<style scoped lang="scss">
.categories{
    .wrap{
        display: flex;
        flex-direction: column;
        gap: 24px;
    }
    &__title{
        text-align: center;
    }
    &__quizzes {
        display: flex;
        justify-content: space-between;
        gap: 16px;
        flex-wrap: wrap;
    }
    &__quiz {
        width: 30%;
    }
}
@media (max-width: 1024px) {
    .categories__quiz{
        width: 50%;
        margin-left: -8px;
        margin-right: -8px;
    }
}
@media (max-width: 600px) {
    .categories__quiz {
        width: 100%;
        margin-left: auto;
        margin-right: auto;
    }
}
</style>