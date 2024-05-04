<script setup lang="ts">
import type {IQuizResults, IQuizGame } from "@/types/quiz.type";
const { t } = useI18n()
const { params } = useRoute()

useHead({
    title: `Quiz id ${params.id}`
})

const info = ref<IQuizGame>(null)
const titleBanner = ref<string>(t('content.main.bannerQuiz'))
const { data } = await useServerFetch<IQuizGame>(`/api/start-quiz/${params.id}`)
info.value = data.value

const onChangeBannerTitle = (event: IQuizResults) => {
    const userObject = JSON.parse(event);
    titleBanner.value = t('content.main.bannerGame', { count: userObject.percentageWins })
}

</script>
<template>
    <MainBanner
     :main-text="t(`${info.name}`)"
     :title=titleBanner
    />
    <QuizGame
     :quizId="info.id"
     :options="info.options"
     @final-game="onChangeBannerTitle"
    />
</template>

<style scoped>

</style>