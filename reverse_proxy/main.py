from fastapi import FastAPI, HTTPException, Request, Response
from pydantic import BaseModel
import random

app = FastAPI()

# Simple in-memory "database"
items = {}

# Request body model
class Item(BaseModel):
    name: str
    description: str | None = None
    price: float
    quantity: int

# GET - read all items
@app.get("/items")
def read_items(request: Request):
    headers = dict(request.headers)
    print("Request Headers:", headers)
    r = random.randint(1, 5)
    if r == 3:
        return Response(content='{"msg": "Random error occurred!"}', status_code=400)
    else:
      return items

@app.get("/get-new-headers")
def read_items():
    return Response(content='Key:Value\nKey2:Value2', status_code=200)

# GET - read a single item
@app.get("/items/{item_id}")
def read_item(item_id: str):
    if item_id not in items:
        raise HTTPException(status_code=404, detail="Item not found")
    return items[item_id]

# POST - create item
@app.post("/items/{item_id}")
def create_item(item_id: str, item: Item):
    if item_id in items:
        raise HTTPException(status_code=400, detail="Item already exists")
    items[item_id] = item
    return {"msg": "Item created", "item": item}

# PUT - update item
@app.put("/items/{item_id}")
def update_item(item_id: str, item: Item):
    if item_id not in items:
        raise HTTPException(status_code=404, detail="Item not found")
    items[item_id] = item
    return {"msg": "Item updated", "item": item}

# DELETE - remove item
@app.delete("/items/{item_id}")
def delete_item(item_id: str):
    if item_id not in items:
        raise HTTPException(status_code=404, detail="Item not found")
    del items[item_id]
    return {"msg": "Item deleted"}
