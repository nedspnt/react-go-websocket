// App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import WebSocketComponent from './WebSocketComponent';
import WebSocketComponent2 from './WebSocketComponent-2';

function App() {
  return (
    <Router>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/page1">Page 1</Link>
            </li>
            <li>
              <Link to="/page2">Page 2</Link>
            </li>
          </ul>
        </nav>
      </div>

      <Routes>
        <Route path="/page1" exact element={<WebSocketComponent />} />
        <Route path="/page2" element={<WebSocketComponent2 />} />
      </Routes>
    </Router>
  );
}

export default App;
