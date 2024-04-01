<script setup lang="ts">

import Logo from 'assets/svg/Logo.svg?component'
</script>
<template>
<header>
    <div class="logo">
        <div class="logo__icon">
            <Logo
                width="40"
                height="40"
            />
        </div>
        <a href="/" class="logo__service-name">Quiz App</a>
    </div>
    <LocaleChanger />
</header>
</template>


<style scoped lang="scss">
 header{
     padding: 16px 40px;
     display: flex;
     justify-content: space-between;
     align-items: center;
 }
 .logo{
     display: flex;
     align-items: center;
     gap: 12px;
     &__icon{
         height: 40px;
     }
     &__service-name{
         color: var(--Content-Primary);
         font-family: Montserrat,sans-serif;
         font-size: 20px;
         font-style: normal;
         font-weight: 700;
         line-height: 24px;
         text-decoration: none;
     }
 }
</style>