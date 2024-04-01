export default defineI18nLocale(locale => {
  const config = useRuntimeConfig()
  const url =  `${config.public.baseUrl}:${config.public.port}/api/translations/${locale}`
  return $fetch(url)
})
