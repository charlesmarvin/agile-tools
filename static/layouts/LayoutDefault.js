export default {
  name: 'LayoutDefault',
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
        <router-view></router-view>
      </main>
      
      <footer class="LayoutDefault__footer">
        made by marvin
      </footer>
    </div>
  `
}
