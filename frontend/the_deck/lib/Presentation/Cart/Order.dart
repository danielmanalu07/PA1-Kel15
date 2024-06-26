import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Controller/ProductController.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:the_deck/Models/Cart_Item.dart';
import 'package:the_deck/Models/Order.dart';
import 'package:the_deck/Presentation/Cart/MyOrder.dart';
import 'package:the_deck/Presentation/Main/main_view.dart';
import 'package:the_deck/Presentation/Profil/profil_view.dart';

class OrderDetailsFormScreen extends StatefulWidget {
  @override
  _OrderDetailsFormScreenState createState() => _OrderDetailsFormScreenState();
}

class _OrderDetailsFormScreenState extends State<OrderDetailsFormScreen> {
  final _formKey = GlobalKey<FormState>();
  final registerController = Get.find<RegisterController>();
  final _productController = Get.put(ProductController());
  final _noteController = TextEditingController();
  String _paymentMethod = 'Cash';
  String _pickUpType = 'Take Away';

  @override
  void initState() {
    super.initState();
    final arguments = Get.arguments;
    if (arguments != null) {
      final List<int> productIds = arguments['products'] ?? [];
      for (int id in productIds) {
        registerController.cartItems.add(CartItem(
          productId: id,
          isChecked: true,
          quantity: 1,
          id: id,
        ));
      }
    }
    _productController.fetchProductList();
  }

  void _submitOrder() async {
    final url = Uri.parse('http://192.168.188.215:8080/order/create');
    final token = registerController.box.read('token');

    final pickUpType = _pickUpType;

    List<int> productIds = registerController.cartItems
        .where((item) => item.isChecked)
        .map((item) => item.productId)
        .toList();

    double total = 0.0;
    registerController.cartItems.forEach((item) {
      if (item.isChecked) {
        final product = _productController.productList
            .firstWhere((prod) => prod.id == item.productId);
        total += product.price * item.quantity;
      }
    });

    final body = json.encode({
      'product_ids': productIds,
      'total': total.toStringAsFixed(2),
      'note': _noteController.text,
      'payment_method': _paymentMethod,
      'table_id': 0, // Set table_id to 0 by default
      'pick_up_type': pickUpType,
    });

    final response = await http.post(
      url,
      headers: {
        'Cookie': 'jwt=$token',
        'Content-Type': 'application/json',
      },
      body: body,
    );

    if (response.statusCode == 200) {
      Get.to(() => ProfilView());
      Get.snackbar('Success', 'Order placed successfully',
          snackPosition: SnackPosition.TOP);
    } else {
      Get.snackbar('Error', 'Failed to place order',
          snackPosition: SnackPosition.TOP);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('Order Details')),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Form(
            key: _formKey,
            child: Column(
              children: [
                TextFormField(
                  controller: _noteController,
                  decoration: InputDecoration(labelText: 'Note'),
                  onSaved: (value) => _noteController.text = value!,
                ),
                DropdownButtonFormField<String>(
                  decoration: InputDecoration(labelText: 'Payment Method'),
                  value: _paymentMethod,
                  items: ['Cash', 'QRIS']
                      .map((method) => DropdownMenuItem(
                            value: method,
                            child: Text(method),
                          ))
                      .toList(),
                  onChanged: (value) => setState(() => _paymentMethod = value!),
                ),
                DropdownButtonFormField<String>(
                  decoration: InputDecoration(labelText: 'Pick Up Type'),
                  value: _pickUpType,
                  items: ['Take Away', 'Dine In']
                      .map((type) => DropdownMenuItem(
                            value: type,
                            child: Text(type),
                          ))
                      .toList(),
                  onChanged: (value) => setState(() => _pickUpType = value!),
                ),
                if (registerController.cartItems.isNotEmpty)
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      SizedBox(height: 20),
                      Text(
                        'Selected Products:',
                        style: TextStyle(
                            fontSize: 18, fontWeight: FontWeight.bold),
                      ),
                      SizedBox(height: 10),
                      Obx(() {
                        return ListView.builder(
                          shrinkWrap: true,
                          itemCount: registerController.cartItems.length,
                          itemBuilder: (context, index) {
                            final cartItem =
                                registerController.cartItems[index];
                            final product = _productController.productList
                                .firstWhere(
                                    (prod) => prod.id == cartItem.productId);
                            if (cartItem.isChecked.obs.isTrue) {
                              return ListTile(
                                title: Text(product.name),
                                subtitle:
                                    Text('Quantity: ${cartItem.quantity}'),
                                trailing: Text(
                                    'Price: Rp ${(cartItem.quantity * product.price).toStringAsFixed(2)}'),
                              );
                            }
                            return SizedBox.shrink();
                          },
                        );
                      }),
                    ],
                  ),
                SizedBox(height: 20),
                ElevatedButton(
                  onPressed: () {
                    if (_formKey.currentState!.validate()) {
                      _formKey.currentState!.save();
                      _submitOrder();
                    }
                  },
                  child: Text('Submit'),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
