import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './ListItems.css'; // New CSS file for styling

const ListItems = ({ onLogout }) => {
  const [items, setItems] = useState([]);
  const [cart, setCart] = useState(null);
  const token = localStorage.getItem('token');

  const fetchUserCart = async () => {
    try {
      const response = await axios.get('http://localhost:8080/carts', {
        headers: {
          Authorization: token,
        },
      });
      const openCart = response.data.carts.find(c => c.status === "open");
      if (openCart) {
        setCart(openCart);
      } else {
        setCart(null);
      }
    } catch (error) {
      console.error('Failed to fetch user cart:', error);
    }
  };

  useEffect(() => {
    const fetchItems = async () => {
      try {
        const response = await axios.get('http://localhost:8080/items');
        setItems(response.data.items);
      } catch (error) {
        console.error('Failed to fetch items:', error);
      }
    };
    fetchItems();
    fetchUserCart();
  }, []);

  const handleAddToCart = async (itemId) => {
    try {
      const response = await axios.post('http://localhost:8080/carts',
        { item_ids: [itemId] },
        {
          headers: {
            Authorization: token,
          },
        }
      );
      setCart(response.data.cart);
      alert(`Item added to cart! Cart ID: ${response.data.cart.id}`);
    } catch (error) {
      console.error('Failed to add item to cart:', error);
      alert('Failed to add item to cart.');
    }
  };

  const handleShowCart = () => {
    if (cart && cart.items && cart.items.length > 0) {
      const cartItems = cart.items.map(item => `Item ID: ${item.id}, Name: ${item.name}`).join('\n');
      window.alert(`Cart ID: ${cart.id}\nCart Items:\n${cartItems}`);
    } else {
      window.alert('Cart is empty.');
    }
  };

  const handleShowOrders = async () => {
    try {
      const response = await axios.get('http://localhost:8080/orders', {
        headers: {
          Authorization: token,
        },
      });
      const orders = response.data.orders;
      if (orders && orders.length > 0) {
        const orderIds = orders.map(order => order.id).join(', ');
        window.alert(`Order IDs: ${orderIds}`);
      } else {
        window.alert('No orders found.');
      }
    } catch (error) {
      console.error('Failed to fetch orders:', error);
      window.alert('Failed to fetch order history.');
    }
  };

  const handleCheckout = async () => {
    if (!cart || cart.status === "ordered" || cart.items.length === 0) {
      window.alert('Your cart is empty or has already been checked out.');
      return;
    }

    try {
      await axios.post('http://localhost:8080/orders',
        { cart_id: cart.id },
        {
          headers: {
            Authorization: token,
          },
        }
      );
      alert('Order successful!');
      setCart(null);
      fetchUserCart();
    } catch (error) {
      console.error('Checkout failed:', error);
      alert('Checkout failed.');
    }
  };

  return (
    <div className="list-items-container">
      <header className="list-items-header">
        <h2 className="title">Available Items</h2>
        <div className="button-group">
          <button onClick={handleShowCart}>Cart ({cart && cart.items ? cart.items.length : 0})</button>
          <button onClick={handleShowOrders}>Order History</button>
          <button onClick={handleCheckout} className="checkout-button">Checkout</button>
        </div>
      </header>
      <div className="items-grid">
        {items.map(item => (
          <div key={item.id} className="item-card">
            <h3>{item.name}</h3>
            <p>Status: <span className="status-badge">{item.status}</span></p>
            <button onClick={() => handleAddToCart(item.id)} className="add-to-cart-button">Add to Cart</button>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ListItems;