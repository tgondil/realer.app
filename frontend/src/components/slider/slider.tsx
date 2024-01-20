import React from 'react';
import "./slider.css";

interface Friend {
  id: number;
  name: string;
}

const Slider: React.FC = () => {
    const friends = [
        { id: 1, name: 'Timothy Edwards' },
        { id: 2, name: 'Mikayla Brown' },
        { id: 3, name: 'Tiffany Calhoun' },
        { id: 4, name: 'Patrick Jones' },
        { id: 5, name: 'Sara Good' },
        { id: 6, name: 'Rachel Becker' },
        { id: 7, name: 'Mrs. Brianna Adams' },
        { id: 8, name: 'Michael Williamson' },
        { id: 9, name: 'William Gray' },
        { id: 10, name: 'Mrs. Ashley Lucas MD' },
        { id: 11, name: 'Christine Morales' },
        { id: 12, name: 'Melissa Smith' },
        { id: 13, name: 'Tyler Horton' },
        { id: 14, name: 'Noah Mccormick' },
        { id: 15, name: 'Alicia Ferrell' },
        { id: 16, name: 'Rose Cruz' },
        { id: 17, name: 'Catherine Camacho' },
        { id: 18, name: 'Megan Jones' },
        { id: 19, name: 'Bradley Ward' },
        { id: 20, name: 'Michael Smith' },
      ];
  return (
    <div className="scrollable-section">
      {friends.map(friend => (
        <div key={friend.id} className="friend-item">
          <p>{friend.name}</p>
        </div>
      ))}
    </div>
  );
};
          

export default Slider;
