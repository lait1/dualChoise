<script setup lang="ts">
import type {ICategoryInfo} from "#build/types/category.type";
const { t } = useI18n()
const props = defineProps<{
    category: ICategoryInfo
}>()

</script>

<template>
    <div class="category-item">
        <div class="category-item__info">
          <p class="category-item__info-name">{{ t(`${props.category.name}`) }}</p>
          <span class="category-item__info-quizzes">{{ props.category.countQuizzes }} quizzes</span>
        </div>
      <img :src="`/images/${props.category.preview}`" :alt="t(`${props.category.name}`)">
    </div>
</template>

<style scoped lang="scss">
.category-item{
  width: 32%;
  height: 126px;
  border-radius: 12px;
  background: var(--Background-Secondary, #FFF);
  padding: 12px;
  border: 1px solid transparent;
  transition: 0.3s;
  box-sizing: border-box;
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  cursor: pointer;
  &:hover{
    border: 1px solid var(--Border-Primary, #180A0A);
    box-shadow: 4px 4px 0 0 rgba(24, 10, 10, 0.24);
  }
  &__info{
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-end;
    &-name{
      margin: 0;
    }
    &-quizzes{
      color: var(--Content-Secondary, #918697);
      font-size: 14px;
      font-weight: 400;
      line-height: 20px
    }
  }
  img{
    width: 102px;
    height: 100%;
    object-fit: cover;
  }
}
@media (max-width: 870px) {
  .category-item {
      width: 50%;
      margin-left: -8px;
      margin-right: -8px;
    }
}
@media (max-width: 600px) {
  .category-item {
    width: 100%;
  }
}
</style>