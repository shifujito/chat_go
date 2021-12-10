/* eslint-disable */
// prettier-ignore
import { AspidaClient, BasicHeaders, dataToURLString } from 'aspida'
// prettier-ignore
import { Methods as Methods0 } from './pet'
// prettier-ignore
import { Methods as Methods1 } from './pet/_petId@number'
// prettier-ignore
import { Methods as Methods2 } from './pet/findByStatus'
// prettier-ignore
import { Methods as Methods3 } from './store/inventory'
// prettier-ignore
import { Methods as Methods4 } from './store/order'
// prettier-ignore
import { Methods as Methods5 } from './store/order/_orderId@number'
// prettier-ignore
import { Methods as Methods6 } from './user'
// prettier-ignore
import { Methods as Methods7 } from './user/_username@string'
// prettier-ignore
import { Methods as Methods8 } from './user/createWithArray'
// prettier-ignore
import { Methods as Methods9 } from './user/createWithList'
// prettier-ignore
import { Methods as Methods10 } from './user/login'
// prettier-ignore
import { Methods as Methods11 } from './users'

// prettier-ignore
const api = <T>({ baseURL, fetch }: AspidaClient<T>) => {
  const prefix = (baseURL === undefined ? 'https://petstore.swagger.io/v2' : baseURL).replace(/\/$/, '')
  const PATH0 = '/pet'
  const PATH1 = '/pet/findByStatus'
  const PATH2 = '/store/inventory'
  const PATH3 = '/store/order'
  const PATH4 = '/user'
  const PATH5 = '/user/createWithArray'
  const PATH6 = '/user/createWithList'
  const PATH7 = '/user/login'
  const PATH8 = '/users'
  const GET = 'GET'
  const POST = 'POST'
  const PUT = 'PUT'
  const DELETE = 'DELETE'

  return {
    pet: {
      _petId: (val1: number) => {
        const prefix1 = `${PATH0}/${val1}`

        return {
          /**
           * Returns a single pet
           * @returns successful operation
           */
          get: (option?: { config?: T }) =>
            fetch<Methods1['get']['resBody'], BasicHeaders, Methods1['get']['status']>(prefix, prefix1, GET, option).json(),
          /**
           * Returns a single pet
           * @returns successful operation
           */
          $get: (option?: { config?: T }) =>
            fetch<Methods1['get']['resBody'], BasicHeaders, Methods1['get']['status']>(prefix, prefix1, GET, option).json().then(r => r.body),
          post: (option: { body: Methods1['post']['reqBody'], config?: T }) =>
            fetch(prefix, prefix1, POST, option, 'URLSearchParams').send(),
          $post: (option: { body: Methods1['post']['reqBody'], config?: T }) =>
            fetch(prefix, prefix1, POST, option, 'URLSearchParams').send().then(r => r.body),
          delete: (option?: { headers?: Methods1['delete']['reqHeaders'], config?: T }) =>
            fetch(prefix, prefix1, DELETE, option).send(),
          $delete: (option?: { headers?: Methods1['delete']['reqHeaders'], config?: T }) =>
            fetch(prefix, prefix1, DELETE, option).send().then(r => r.body),
          $path: () => `${prefix}${prefix1}`
        }
      },
      findByStatus: {
        /**
         * Multiple status values can be provided with comma separated strings
         * @returns successful operation
         */
        get: (option: { query: Methods2['get']['query'], config?: T }) =>
          fetch<Methods2['get']['resBody'], BasicHeaders, Methods2['get']['status']>(prefix, PATH1, GET, option).json(),
        /**
         * Multiple status values can be provided with comma separated strings
         * @returns successful operation
         */
        $get: (option: { query: Methods2['get']['query'], config?: T }) =>
          fetch<Methods2['get']['resBody'], BasicHeaders, Methods2['get']['status']>(prefix, PATH1, GET, option).json().then(r => r.body),
        $path: (option?: { method?: 'get'; query: Methods2['get']['query'] }) =>
          `${prefix}${PATH1}${option && option.query ? `?${dataToURLString(option.query)}` : ''}`
      },
      /**
       * @param option.body - Pet object that needs to be added to the store
       */
      post: (option: { body: Methods0['post']['reqBody'], config?: T }) =>
        fetch(prefix, PATH0, POST, option).send(),
      /**
       * @param option.body - Pet object that needs to be added to the store
       */
      $post: (option: { body: Methods0['post']['reqBody'], config?: T }) =>
        fetch(prefix, PATH0, POST, option).send().then(r => r.body),
      /**
       * @param option.body - Pet object that needs to be added to the store
       */
      put: (option: { body: Methods0['put']['reqBody'], config?: T }) =>
        fetch(prefix, PATH0, PUT, option).send(),
      /**
       * @param option.body - Pet object that needs to be added to the store
       */
      $put: (option: { body: Methods0['put']['reqBody'], config?: T }) =>
        fetch(prefix, PATH0, PUT, option).send().then(r => r.body),
      $path: () => `${prefix}${PATH0}`
    },
    store: {
      inventory: {
        /**
         * Returns a map of status codes to quantities
         * @returns successful operation
         */
        get: (option?: { config?: T }) =>
          fetch<Methods3['get']['resBody'], BasicHeaders, Methods3['get']['status']>(prefix, PATH2, GET, option).json(),
        /**
         * Returns a map of status codes to quantities
         * @returns successful operation
         */
        $get: (option?: { config?: T }) =>
          fetch<Methods3['get']['resBody'], BasicHeaders, Methods3['get']['status']>(prefix, PATH2, GET, option).json().then(r => r.body),
        $path: () => `${prefix}${PATH2}`
      },
      order: {
        _orderId: (val2: number) => {
          const prefix2 = `${PATH3}/${val2}`

          return {
            /**
             * For valid response try integer IDs with value >= 1 and <= 10. Other values will generated exceptions
             * @returns successful operation
             */
            get: (option?: { config?: T }) =>
              fetch<Methods5['get']['resBody'], BasicHeaders, Methods5['get']['status']>(prefix, prefix2, GET, option).json(),
            /**
             * For valid response try integer IDs with value >= 1 and <= 10. Other values will generated exceptions
             * @returns successful operation
             */
            $get: (option?: { config?: T }) =>
              fetch<Methods5['get']['resBody'], BasicHeaders, Methods5['get']['status']>(prefix, prefix2, GET, option).json().then(r => r.body),
            /**
             * For valid response try integer IDs with positive integer value. Negative or non-integer values will generate API errors
             */
            delete: (option?: { config?: T }) =>
              fetch(prefix, prefix2, DELETE, option).send(),
            /**
             * For valid response try integer IDs with positive integer value. Negative or non-integer values will generate API errors
             */
            $delete: (option?: { config?: T }) =>
              fetch(prefix, prefix2, DELETE, option).send().then(r => r.body),
            $path: () => `${prefix}${prefix2}`
          }
        },
        /**
         * @param option.body - order placed for purchasing the pet
         * @returns successful operation
         */
        post: (option: { body: Methods4['post']['reqBody'], config?: T }) =>
          fetch<Methods4['post']['resBody'], BasicHeaders, Methods4['post']['status']>(prefix, PATH3, POST, option).json(),
        /**
         * @param option.body - order placed for purchasing the pet
         * @returns successful operation
         */
        $post: (option: { body: Methods4['post']['reqBody'], config?: T }) =>
          fetch<Methods4['post']['resBody'], BasicHeaders, Methods4['post']['status']>(prefix, PATH3, POST, option).json().then(r => r.body),
        $path: () => `${prefix}${PATH3}`
      }
    },
    user: {
      _username: (val1: string) => {
        const prefix1 = `${PATH4}/${val1}`

        return {
          /**
           * @returns successful operation
           */
          get: (option?: { config?: T }) =>
            fetch<Methods7['get']['resBody'], BasicHeaders, Methods7['get']['status']>(prefix, prefix1, GET, option).json(),
          /**
           * @returns successful operation
           */
          $get: (option?: { config?: T }) =>
            fetch<Methods7['get']['resBody'], BasicHeaders, Methods7['get']['status']>(prefix, prefix1, GET, option).json().then(r => r.body),
          /**
           * This can only be done by the logged in user.
           * @param option.body - Updated user object
           */
          put: (option: { body: Methods7['put']['reqBody'], config?: T }) =>
            fetch(prefix, prefix1, PUT, option).send(),
          /**
           * This can only be done by the logged in user.
           * @param option.body - Updated user object
           */
          $put: (option: { body: Methods7['put']['reqBody'], config?: T }) =>
            fetch(prefix, prefix1, PUT, option).send().then(r => r.body),
          /**
           * This can only be done by the logged in user.
           */
          delete: (option?: { config?: T }) =>
            fetch(prefix, prefix1, DELETE, option).send(),
          /**
           * This can only be done by the logged in user.
           */
          $delete: (option?: { config?: T }) =>
            fetch(prefix, prefix1, DELETE, option).send().then(r => r.body),
          $path: () => `${prefix}${prefix1}`
        }
      },
      createWithArray: {
        /**
         * @param option.body - List of user object
         */
        post: (option: { body: Methods8['post']['reqBody'], config?: T }) =>
          fetch(prefix, PATH5, POST, option).send(),
        /**
         * @param option.body - List of user object
         */
        $post: (option: { body: Methods8['post']['reqBody'], config?: T }) =>
          fetch(prefix, PATH5, POST, option).send().then(r => r.body),
        $path: () => `${prefix}${PATH5}`
      },
      createWithList: {
        /**
         * @param option.body - List of user object
         */
        post: (option: { body: Methods9['post']['reqBody'], config?: T }) =>
          fetch(prefix, PATH6, POST, option).send(),
        /**
         * @param option.body - List of user object
         */
        $post: (option: { body: Methods9['post']['reqBody'], config?: T }) =>
          fetch(prefix, PATH6, POST, option).send().then(r => r.body),
        $path: () => `${prefix}${PATH6}`
      },
      login: {
        /**
         * @returns successful operation
         */
        get: (option: { query: Methods10['get']['query'], config?: T }) =>
          fetch<Methods10['get']['resBody'], Methods10['get']['resHeaders'], Methods10['get']['status']>(prefix, PATH7, GET, option).text(),
        /**
         * @returns successful operation
         */
        $get: (option: { query: Methods10['get']['query'], config?: T }) =>
          fetch<Methods10['get']['resBody'], Methods10['get']['resHeaders'], Methods10['get']['status']>(prefix, PATH7, GET, option).text().then(r => r.body),
        $path: (option?: { method?: 'get'; query: Methods10['get']['query'] }) =>
          `${prefix}${PATH7}${option && option.query ? `?${dataToURLString(option.query)}` : ''}`
      },
      /**
       * This can only be done by the logged in user.
       * @param option.body - Created user object
       */
      post: (option: { body: Methods6['post']['reqBody'], config?: T }) =>
        fetch(prefix, PATH4, POST, option).send(),
      /**
       * This can only be done by the logged in user.
       * @param option.body - Created user object
       */
      $post: (option: { body: Methods6['post']['reqBody'], config?: T }) =>
        fetch(prefix, PATH4, POST, option).send().then(r => r.body),
      $path: () => `${prefix}${PATH4}`
    },
    users: {
      /**
       * @returns successful operation
       */
      get: (option?: { config?: T }) =>
        fetch<Methods11['get']['resBody'], BasicHeaders, Methods11['get']['status']>(prefix, PATH8, GET, option).json(),
      /**
       * @returns successful operation
       */
      $get: (option?: { config?: T }) =>
        fetch<Methods11['get']['resBody'], BasicHeaders, Methods11['get']['status']>(prefix, PATH8, GET, option).json().then(r => r.body),
      $path: () => `${prefix}${PATH8}`
    }
  }
}

// prettier-ignore
export type ApiInstance = ReturnType<typeof api>
// prettier-ignore
export default api
