import Alerts from '../components/Alerts.js'
export default {
  name: 'LayoutDefault',
  components: {
    Alerts
  },
  template: `
    <div class="LayoutDefault">
      <nav class="LayoutDefault__nav">
        <div class="col">
          <h1 title="Agile Tools" class="home">
            <router-link to="/">A-T</router-link>
          </h1>
        </div>
      </nav>

      <main class="LayoutDefault__main">
        <alerts></alerts>
        <router-view></router-view>
      </main>
      
      <footer class="LayoutDefault__footer">
        made by marvin
      </footer>
    </div>
  `
}
