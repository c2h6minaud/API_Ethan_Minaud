import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [artists, setArtists] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/artists")
      .then((res) => res.json())
      .then((data) => setArtists(data));
  }, []);

  return (
    <div>
      <header className="header">
        <h1>Artistes</h1>
      </header>

      <div className="grid">
        {artists.map((artist) => (
          <div className="card" key={artist.id}>
            <h2>{artist.name}</h2>
            <p>{artist.genre}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;