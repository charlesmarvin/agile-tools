const Alerts = () => import('./Alerts.js')
const Board = () => import('./Board.js')
const CreateBoardForm = () => import('./CreateBoardForm.js')

export default {
  name: 'App',
  components: {
    Alerts,
    Board,
    CreateBoardForm
  },
  template: `
  <div class="container">
    <div class="flex-grid">
      <div class="col">
        <h1 title="Agile Tools" class="home">
          <a href="/">A-T</a>
        </h1>
      </div>
      <div class="flex-grid">
          <Alerts></Alerts> 
      </div>
    </div>

    <div class="flex-grid">
      <div class="col">
        <CreateBoardForm></CreateBoardForm>
      </div>
    </div>

    <div class="flex-grid">
      <div class="col">
          <Board></Board>
      </div>
    </div>
  </div>
  `
}
