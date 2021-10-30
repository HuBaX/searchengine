import './App.css';
import Header from './components/header';
import SearchRequest from './components/searchrequest';
import background from './background.png'



function App() {
  return (
    <div  style={{ backgroundImage:`url(${background})`,  backgroundSize: 'cover', backgroundRepeat: 'no-repeat'}}>
      <Header></Header>
      <SearchRequest></SearchRequest>
    </div>
  );
}


export default App;
