<script setup lang="ts" xmlns="http://www.w3.org/1999/html">
const {t} = useI18n()
import Arrow from 'assets/svg/Arrow.svg?component'

const props = defineProps<{
    mainText: string,
    title: string
}>()
</script>

<template>
    <section class="main-banner wrap">
        <Arrow class="main-banner__arrow main-banner__arrow-left"/>
        <div class="main-banner__content">
            <h1>{{ props.mainText }}</h1>
            <h3>{{ props.title }}</h3>
        </div>
        <Arrow class="main-banner__arrow main-banner__arrow-right"/>
    </section>
</template>

<style scoped lang="scss">
.main-banner {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;

  &__content {
    text-align: center;
    align-self: flex-start;

    h3 {
      margin: 16px 0;
    }
  }

  &__arrow {
    overflow: unset;
    width: 150px;
    margin-bottom: 10px;

    &-right {
      transform: scale(-1, 1)
    }
  }
}
@media (max-width: 1024px) {
    .main-banner__arrow {
        width: 100px;
    }
}
@media (max-width: 500px) {
    .main-banner__arrow {
        display: none;
    }
    .main-banner{
        justify-content: center;
    }
}
</style>