type useFetchType = typeof useFetch

export const useServerFetch: useFetchType = (path, options = {}) => {
  options.baseURL = `http://app:4000`;
  return useFetch(path, options)
}