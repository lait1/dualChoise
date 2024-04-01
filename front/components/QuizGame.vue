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

const selectFirstOption = (option: IQuizQuestion) => {
    playClickSound()
    firstQuiz.value.$el.classList.add('quiz-game__winner')
    setTimeout(() => {
        setFavouriteOption(option)
        firstQuiz.value.$el.classList.remove('quiz-game__winner')
        secondChoice.value = questions.value[0]
        questions.value = questions.value.slice(1)
    }, 700)
}

const selectSecondOption = (option: IQuizQuestion) => {
    playClickSound()
    secondQuiz.value.$el.classList.add('quiz-game__winner')
    setTimeout(() => {
        setFavouriteOption(option)
        secondQuiz.value.$el.classList.remove('quiz-game__winner')
        firstChoice.value = questions.value[0]
        questions.value = questions.value.slice(1)
    }, 700)
}

const setFavouriteOption = async (option: IQuizQuestion) => {
  favouriteChoice.value = option
  if (questions.value.length === 0) {
      const { data } = await useAPIFetch<IQuizResults>("/api/send-result", {
          method: 'POST',
          body: {quizId: props.quizId, optionId: favouriteChoice.value.id}
      });

      emit('final-game', data.value)

      showResult.value = true
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
                @click="selectFirstOption(firstChoice)"
            />
            <OR/>
            <Quiz
                ref="secondQuiz"
                :size="240"
                :quiz="secondChoice"
                @click="selectSecondOption(secondChoice)"
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