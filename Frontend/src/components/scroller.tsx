import InfiniteScroll from "react-infinite-scroll-component";
import { useState } from "react";
import RecipeCard from "./recipecard";

interface RecipePreview {id: number, minutes : number, name: string , description: string}

function Scroller() {
  const [items, setItems] = useState<RecipePreview[]>([]);

  const [hasMore, sethasMore] = useState(true);

  const [page, setpage] = useState(2);

  const fetchComments = async () => {
    const res = await fetch(
      `https://jsonplaceholder.typicode.com/comments?_page=${page}&_limit=20`
      // For json server use url below
      // `http://localhost:3004/comments?_page=${page}&_limit=20`
    );
    const data = await res.json();
    return data;
  };

  const fetchData = async () => {
    const commentsFormServer = await fetchComments();

    setItems([...items, ...commentsFormServer]);
    if (commentsFormServer.length === 0 || commentsFormServer.length < 20) {
      sethasMore(false);
    }
    setpage(page + 1);
  };


  return (
    <InfiniteScroll
        dataLength={items.length} //This is important field to render the next data
        next={fetchData}
        hasMore={hasMore}
        loader={<div></div>}
        endMessage={
            <p style={{ textAlign: 'center' }}>
            <b>Yay! You have seen it all</b>
            </p>
        }
    >
      <div className="container">
        <div className="row m-2">
          {items.map((item) => {
            return <RecipeCard recipeName={item.name} description={item.description} time={item.minutes} />;
          })}
        </div>
      </div>
    </InfiniteScroll>
  );
}

export default Scroller;

