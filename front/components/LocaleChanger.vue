<script setup lang="ts">
const { locale, locales, setLocale } = useI18n()
</script>

<template>
    <select
        v-model="locale"
        @change="setLocale($event.target.value)"
    >
        <option
            v-for="locale in locales"
            :key="locale.code"
            :value="locale.code"
        >
         {{ locale.name }}
        </option>
    </select>
</template>


<style scoped>

</style>