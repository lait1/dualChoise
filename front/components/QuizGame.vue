<script setup lang="ts">
import type {IQuizQuestion, IQuizResults} from "@types/quiz.type";
import OR from 'assets/svg/OR.svg?component'

const emit = defineEmits<{
    (e: 'final-game', value: IQuizResults): void,
}>()

const props = defineProps<{
  quizId: number,
  options: IQuizQuestion[],
}>()

const questions = ref<IQuizQuestion[]>(props.options)
const favouriteChoice = ref<IQuizQuestion>(null)

const firstChoice = ref<IQuizQuestion>(null)
const firstQuiz = ref<HTMLElement>()

const secondChoice = ref<IQuizQuestion>(null)
const secondQuiz = ref<HTMLElement>()

const showResult = ref<boolean>(false)
const isClickAllowed = ref<boolean>(true)

const selectOption = (option: IQuizQuestion, quizRef: Ref<HTMLElement>) => {
  try {
    if (!isClickAllowed.value) return
    isClickAllowed.value = false
    playClickSound()
    quizRef.$el.classList.add('quiz-game__winner')

    setTimeout(() => {
      sendFavouriteOption(option)
      quizRef.$el.classList.remove('quiz-game__winner')

      if (option === firstChoice.value) {
        secondChoice.value = questions.value[0]
      } else {
        firstChoice.value = questions.value[0]
      }

      questions.value = questions.value.slice(1)
      isClickAllowed.value = true
    }, 700)
  } catch (e) {
    isClickAllowed.value = true
    console.error(e)
  }
}

const sendFavouriteOption = async (option: IQuizQuestion) => {
  favouriteChoice.value = option
  if (showResult.value === false && questions.value.length === 0) {
      showResult.value = true

      const { data } = await useAPIFetch<IQuizResults>("/api/send-result", {
          method: 'POST',
          body: {quizId: props.quizId, optionId: favouriteChoice.value.id}
      });

      emit('final-game', data.value)
      return
  }
}

const playClickSound = () => {
  const audio = new Audio('/sounds/click.mp3')
  audio.play()
}
const start = () => {
  questions.value.sort((a: IQuizQuestion, b: IQuizQuestion) => b.priority - a.priority);
  [firstChoice.value, secondChoice.value] = [questions.value[0], questions.value[1]]
  questions.value = questions.value.slice(2)
}
start()
</script>
<template>
    <section class="quiz-game">
        <QuizWinner
            v-if="showResult"
            :winner="favouriteChoice"
        />
        <div v-else class="quiz-game__container wrap">
            <Quiz
                ref="firstQuiz"
                :size="240"
                :quiz="firstChoice"
                @click="isClickAllowed && selectOption(firstChoice, firstQuiz)"
            />
            <OR/>
            <Quiz
                ref="secondQuiz"
                :size="240"
                :quiz="secondChoice"
                @click="isClickAllowed && selectOption(secondChoice, secondQuiz)"
            />
        </div>
    </section>
</template>

<style lang="scss">
.quiz-game__winner {
  border: 1px solid var(--Border-Success, #5AC386);
  box-shadow: 4px 4px 0 0 rgba(90, 195, 134, 0.24);
  cursor: not-allowed;

  &:hover {
    border: 1px solid var(--Border-Success, #5AC386);
    box-shadow: 4px 4px 0 0 rgba(90, 195, 134, 0.24);
  }

  .quiz__checkmark {
    display: block;
  }
}

.quiz-game__container {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 16px;
  padding-top: 0;
    @media (max-width: 1024px) {
        flex-direction: column;
    }
}
</style>