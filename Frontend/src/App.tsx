import './App.css';
import Header from './components/header';
import SearchRequest from './components/searchrequest';
import background from './background.png'



function App() {
  // <div style={{position: 'absolute', marginLeft: 'auto', marginRight: 'auto', left: 0, right: 0, textAlign: 'center'}}>
  return (
    <div  style={{ backgroundImage:`url(${background})`,  backgroundSize: 'cover', backgroundAttachment: 'fixed', minHeight: '100vh', width: '100'}}>
      <div>
        <Header></Header>
        <SearchRequest></SearchRequest>
      </div>
    </div>
  );
}


export default App;
