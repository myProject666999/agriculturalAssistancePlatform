import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import { Button, Tabbar, TabbarItem, NavBar, Toast, Swiper, SwiperItem, Search, Grid, GridItem, Card, Stepper, SubmitBar, List, PullRefresh, ActionSheet, Popup, Cell, CellGroup, Field, Divider, Empty, Loading, Overlay, Picker, DatetimePicker, Area, RadioGroup, Radio, Checkbox, CheckboxGroup, Tag, Badge, Progress, CountDown, NoticeBar } from 'vant'
import 'vant/lib/index.css'
import './styles/index.scss'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.use(Button)
app.use(Tabbar)
app.use(TabbarItem)
app.use(NavBar)
app.use(Toast)
app.use(Swiper)
app.use(SwiperItem)
app.use(Search)
app.use(Grid)
app.use(GridItem)
app.use(Card)
app.use(Stepper)
app.use(SubmitBar)
app.use(List)
app.use(PullRefresh)
app.use(ActionSheet)
app.use(Popup)
app.use(Cell)
app.use(CellGroup)
app.use(Field)
app.use(Divider)
app.use(Empty)
app.use(Loading)
app.use(Overlay)
app.use(Picker)
app.use(DatetimePicker)
app.use(Area)
app.use(RadioGroup)
app.use(Radio)
app.use(Checkbox)
app.use(CheckboxGroup)
app.use(Tag)
app.use(Badge)
app.use(Progress)
app.use(CountDown)
app.use(NoticeBar)

app.mount('#app')
