import React from 'react'
import CardItem from './CardItem'
import './Cards.css'

function Cards() {
  return (
    <div className='cards'>
        <h1>CheckOut the Epic Destination</h1>
        <div className="cards__container">
            <div className="cards__wrapper">
                <ul className="cards__items">
                    <CardItem
                        src='IandV\img2.jpg'
                        text='Explore The Hidden Waters'
                        label='Adventure'
                        path='/services'
                    />
                </ul>
            </div>
        </div>
    </div>
  )
}

export default Cards    