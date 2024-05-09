<script setup lang="ts">
import type {ICategoryInfo} from "#build/types/category.type";
import type {IQuiz} from "#build/types/quiz.type";
import ArrowRight from 'assets/svg/Arrow_Right.svg?component'

const { t } = useI18n()
// const limit = isDesktop ? 9 : 6
const limit = 6
const categories = ref<ICategoryInfo[]>([])
const { data, error } = await useServerFetch<IQuiz[]>("/api/get-categories", { query: { limit: limit } })
categories.value = data.value
</script>

<template>
    <section class="popular-categories">
        <div class="wrap">
          <Header :text="t('content.main.categories')" />
          <Categories v-if="categories.length > 0" :categories="categories" />
            <div class="popular-categories__control">
              <ActionButton
                      text="View all categories"
                      @click="navigateTo('/categories',{ external: true })"
              >
                  <template #icon>
                    <ArrowRight class="svg"/>
                  </template>
              </ActionButton>
            </div>
        </div>
    </section>
</template>


<style scoped lang="scss">
.popular-categories{
  .wrap{
    display: flex;
    flex-direction: column;
    gap: 24px;
  }
  &__control {
    display: flex;
    justify-content: center;
  }
}
.svg {
  margin-left: 8px;
}
</style>