import {createI18n} from "vue-i18n"
import deDE from '../locales/de-DE.json'

export const i18n = createI18n({
  locale: 'de-DE',
  messages: {
    'de-DE': deDE,
  }
})

