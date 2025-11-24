import api from "./api";

// CREATE
export const adminCreatePizza = ({ name, price, typeId, photo }) =>
  api.post("/admin/pizzas/create", {
    Name: name,
    Price: Number(price),
    TypeId: Number(typeId),
    Photo: photo,
  });

// POST /admin/pizzas/create/type
export const adminCreateType = ({ name }) =>
  api.post("/admin/pizzas/create/type", {
    Name: name,
  });


// UPDATE  (your backend uses PUT /admin/pizzas/update)
export const adminUpdatePizza = ({ id, typeId, name, price, photo }) =>
  api.put("/admin/pizzas/update", {
    Id: Number(id),
    TypeId: Number(typeId),
    Name: name,
    Price: Number(price),
    Photo: photo,
  });

// DELETE pizza
export const adminDeletePizza = (id) =>
  api.delete(`/admin/pizzas/delete/${id}`);