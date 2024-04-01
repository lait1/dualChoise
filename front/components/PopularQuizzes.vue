<script setup lang="ts">
import type {IQuiz} from "#build/types/quiz.type.js";
const { isDesktop, isMobile, isTablet } = useDevice();
useSeoMeta({
    title: 'Best choice',
    ogTitle: 'My Amazing Site',
    description: 'This is my amazing site, let me tell you all about it.',
    ogDescription: 'This is my amazing site, let me tell you all about it.',
    ogImage: 'https://example.com/image.png',
})
const {t} = useI18n()
// const limit = isDesktop ? 3 : 4
const limit = 4
const imageSize = isDesktop ? 200 : 250
const tipQuizzes = ref<IQuiz[]>([])
const {data, pending} = await useAPIFetch<IQuiz[]>("/api/get-popular-quizzes", { query: { limit: limit } })
tipQuizzes.value = data.value
</script>

<template>
    <h1 v-if="pending">i am pending</h1>
    <section class="popular-quizzes">
        <div class="wrap">
            <Header :text="t('content.main.popularQuizzes')" />
            <div class="popular-quizzes__wrap">
                <Quiz
                    class="popular-quizzes__quiz"
                    v-for="quiz in tipQuizzes"
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
.star {
  width: 32px;
  height: 32px;
}

.popular-quizzes {
  background: var(--Background-Accent);
  position: relative;

  .wrap{
    display: flex;
    flex-direction: column;
    gap: 24px;
  }

  &__quiz {
    width: 33.3333%;
  }

  &__wrap {
    display: flex;
    gap: 16px;
  }

  &::before {
    top: -5px;
  }

  &::after {
    bottom: -5px;
    rotate: 180deg;
  }
}

.popular-quizzes::before, .popular-quizzes::after {
  content: "";
  position: absolute;
  width: 100%;
  height: 20px;
  background-image: url("assets/svg/Border.svg");
  background-repeat: repeat-x;
  z-index: 1;
}
@media (max-width: 1024px) {
    .popular-quizzes{
        &__wrap{
            flex-wrap: wrap;
            justify-content: space-between;
        }
        &__quiz {
            width: 50%;
            margin-left: -8px;
            margin-right: -8px;
        }
    }
}
@media (max-width: 600px) {
    .popular-quizzes__quiz {
        width: 100%;
        margin-left: auto;
        margin-right: auto;
    }
}
</style>