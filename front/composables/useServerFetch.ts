type useFetchType = typeof useFetch

export const useServerFetch: useFetchType = (path, options = {}) => {
  const config = useRuntimeConfig()

  options.baseURL = `${config.public.serverUrl}`;
  return useFetch(path, options)
}