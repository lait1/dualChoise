export default defineI18nLocale(locale => {
  const config = useRuntimeConfig()
  const url =  `${config.public.baseUrl}/api/translations/${locale}`

  return $fetch(url)
})
