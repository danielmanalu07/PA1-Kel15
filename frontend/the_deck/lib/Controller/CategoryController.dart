import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

import 'package:the_deck/Models/Category.dart';

class CategoryController extends GetxController {
  var categories = <Category>[].obs;
  var isLoading = true.obs;

  @override
  void onInit() {
    fetchCategories();
    super.onInit();
  }

  void fetchCategories() async {
    try {
      isLoading(true);
      final response =
          await http.get(Uri.parse('http://192.168.30.215:8080/category'));
      if (response.statusCode == 200) {
        List<dynamic> data = jsonDecode(response.body)['message'];
        categories.value = data.map((json) => Category.fromJson(json)).toList();
      } else {
        // Handle error
        Get.snackbar('Error', 'Failed to fetch categories');
      }
    } catch (e) {
      // Handle error
      Get.snackbar('Error', 'Failed to fetch categories');
    } finally {
      isLoading(false);
    }
  }
}