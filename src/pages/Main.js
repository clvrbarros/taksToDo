import React, { useState } from 'react';
import './Main.css'

export default function Main() {

  const [title, setTitle] = useState('');
  const [desc, setDesc] = useState('');
  const [value, setValue] = useState(3);

  function handleSubmit (e) {
    e.preventDefault();
    console.log(title, desc, value)
  }

  return (
    <div className="main-container">
        <h1> Welcome! </h1>
        <p>Are you want add tasks?</p>
        <div className="add-task">
          <form onSubmit={handleSubmit}>
              <p>Describe your task and choice the priority.</p>
              <input type="text" 
                     name="title"
                     value={title}
                     placeholder="Title"
                     onChange={e => setTitle(e.target.value)}
              />
              <textarea id="description" 
                        cols="30" rows="5"
                        value={desc}
                        placeholder="Description"
                        onChange={e => setDesc(e.target.value)}
              />
            <select onChange={e => setValue(e.target.value)}>
              <option value="3">Low</option>
              <option value="2">Medium</option>
              <option value="1">Hight</option>
            </select>

            <button type="submit">Submit</button>
          </form>

        </div>
    </div>
  );
}
