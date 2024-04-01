<script setup lang="ts">
const props = withDefaults(defineProps<{
  text: string,
  disabled?: boolean,
}>(), {
  disabled: false,
})
</script>
<template>
    <button
     class="action-button"
     type="button"
     :disabled="props.disabled"
    >
        {{ props.text }}
        <slot name="icon" />
    </button>
</template>

<style scoped lang="scss">
button {
  padding: 12px 24px;
  border-radius: 12px;
  border: 1px solid var(--Content-Primary, #180A0A);
  background: var(--Background-Accent, #CECEFF);
  box-shadow: 3px 3px 0 0 #180A0A;
  font-size: 16px;
  font-style: normal;
  font-weight: 600;
  line-height: 22px;
  cursor: pointer;
  transition: 0.3s;
  display: flex;
  justify-content: space-between;
  &:hover {
    background: var(--Background-Accent_hover, #BEBEFF);
  }

  &:active {
    background: var(--Background-Accent_active, #AFAFFF);
  }

  &:disabled,
  &[disabled] {
    background: var(--Background-Accent_disabled, #CECEFF);
    cursor: not-allowed;
  }
}

</style>