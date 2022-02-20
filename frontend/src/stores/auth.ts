import {defineStore} from 'pinia'
import jwtDecode, {JwtPayload} from 'jwt-decode';

type AuthState = {
  user: string
  loggedIn: boolean
}

export const useAuth = defineStore<string, AuthState>('auth', {
  state: () => {
    let user: string = ''

    try {
      const cookie = document.cookie.match(/(^|;)\s*kuefa\s*=\s*([^;]+)/)?.pop() || ''
      const claims = jwtDecode<JwtPayload>(cookie)
      user = claims.sub || ''
    } catch (e) {
      console.error(e)
    }

    return {
      user: user,
      loggedIn: user !== '',
    }
  }
})

