import Vue from 'vue'
import { 
  Button, 
  Select,
  Option,
  OptionGroup,
  Input,
  Icon,
  Form, 
  FormItem,
  Message,
  Container,
  Aside,
  Header,
  Main,
  Footer,
  Menu,
  Submenu,
  MenuItem,
  Card,
  Row,
  Col,
  Table,
  TableColumn,
  Pagination,
  MessageBox,
  Dialog,
  Radio,
  RadioGroup,
  RadioButton,
  Upload,
  Image,
  Dropdown,
  DropdownMenu,
  DropdownItem,
  Link,
  Tooltip,
  Divider,
  Tabs,
  TabPane,
  Tag,
  ColorPicker,
  Autocomplete,
  MenuItemGroup} from 'element-ui'


import 'element-ui/lib/theme-chalk/index.css';
import 'element-ui/lib/theme-chalk/display.css';

Vue.prototype.$message = Message
//MessageBox不需要use
Vue.prototype.$confirm = MessageBox.confirm

// 全局注册插件
Vue.use(Button)
Vue.use(Select)
Vue.use(Input)
Vue.use(Icon)
Vue.use(Form)
Vue.use(FormItem)
Vue.use(Container)
Vue.use(Header)
Vue.use(Aside)
Vue.use(Main)
Vue.use(Footer)

Vue.use(Menu)
Vue.use(Submenu)
Vue.use(MenuItem)
Vue.use(MenuItemGroup)

Vue.use(Card)
Vue.use(Row)
Vue.use(Col)

Vue.use(Option)
Vue.use(OptionGroup)

Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Pagination)
Vue.use(Dialog)

Vue.use(Radio)
Vue.use(RadioGroup)
Vue.use(RadioButton)
Vue.use(Upload)
Vue.use(Image)

Vue.use(Dropdown)
Vue.use(DropdownMenu)
Vue.use(DropdownItem)

Vue.use(Link)
Vue.use(Tooltip)
Vue.use(Divider)

Vue.use(Tabs)
Vue.use(TabPane)
Vue.use(Tag)

Vue.use(ColorPicker)

Vue.use(Autocomplete)