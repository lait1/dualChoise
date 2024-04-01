<script setup lang="ts">
import type {IQuiz, IQuizQuestion} from "#build/types/quiz.type";
import Checkmark from 'assets/svg/Checkmark.svg?component'

const { t } = useI18n()

const props = withDefaults(defineProps<{
    size?: number,
    quiz: IQuiz | IQuizQuestion
}>(), {
    size: 200
})
</script>
<template>
    <transition name="image" mode="out-in">
     <div :key="props.quiz" class="quiz">
          <img
            :src="`/images/${props.quiz.preview}`"
            :alt="t(`${props.quiz.name}`)"
            :height="props.size"
          >
         <p class="quiz__name">{{ t(`${props.quiz.name}`) }}</p>
         <Checkmark class="quiz__checkmark" />
     </div>
    </transition>
</template>

<style lang="scss">
.image-enter-active,
.image-leave-active {
  transition: opacity .5s;
}

.image-enter,
.image-leave-to {
  opacity: 0;
}
.quiz {
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-around;
    gap: 16px;
    border-radius: 12px;
    background: var(--Background-Secondary, #FFF);
    padding: 8px 8px 16px 8px;
    border: 1px solid transparent;
    box-sizing: border-box;
    cursor: pointer;
    &__checkmark{
        display: none;
        position: absolute;
        top: 16px;
        left: 16px;
    }
    img {
        width: 100%;
        object-fit: cover;
        border-radius: 8px;
    }
    &__name{
        margin: 0;
    }
    &:hover{
        border: 1px solid var(--Border-Primary, #180A0A);
        box-shadow: 4px 4px 0 0 rgba(24, 10, 10, 0.24);
    }
}
</style>