<script setup lang="ts">
import type {IQuizQuestion} from "#build/types/quiz.type.js";
import Crown from 'assets/svg/Crown.svg?component'
const { isDesktop } = useDevice();
const props = defineProps<{
    winner: IQuizQuestion,
}>()

const audio = new Audio('/sounds/congracs.mp3')
audio.play()
const imageSize = isDesktop ? 200 : 250
</script>

<template>
    <div class="quiz-winner__container wrap">
     <Crown />
     <Quiz
         class="quiz-winner"
        :quiz="props.winner"
         :size="imageSize"
     />
     <ActionButton
         text="To the main page"
         @click="navigateTo('/', { external: true })"
     />
    </div>
</template>

<style scoped>
.quiz-winner__container {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 24px;
    padding-top: 0;
}
.quiz-winner{
    &:hover{
        border: 1px solid transparent;
        box-shadow: none;
        cursor: auto;
    }
}
@media (max-width: 600px) {
    .quiz-winner {
        width: 100%;
        margin-left: auto;
        margin-right: auto;
    }
}
</style>