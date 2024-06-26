import 'package:get/get.dart';
import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/assets_constantes.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Presentation/Auth/screens/read_only.dart';
import 'package:the_deck/Presentation/Base/base.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:gap/gap.dart';

class PersonalDataView extends StatefulWidget {
  const PersonalDataView({Key? key}) : super(key: key);

  @override
  State<PersonalDataView> createState() => _PersonalDataViewState();
}

class _PersonalDataViewState extends State<PersonalDataView> {
  final RegisterController _controller = Get.put(RegisterController());

  @override
  void initState() {
    super.initState();
    _controller.getUserProfile();
  }

  @override
  Widget build(BuildContext context) {
    return Obx(() {
      final customer = _controller.userProfile.value;
      final isLoggedIn = customer != null;

      return Scaffold(
        appBar: buildAppBar(
          buildContext: context,
          screenTitle: "Personal Data",
        ),
        body: Padding(
          padding: EdgeInsets.symmetric(horizontal: getWidth(24)),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              const Gap(24),
              Stack(
                children: [
                  CircleAvatar(
                    backgroundImage: isLoggedIn
                        ? NetworkImage(
                            'http://192.168.188.215:8080/customer/image/${customer?.image}')
                        : AssetImage('assets/images/user_icon.png')
                            as ImageProvider,
                    radius: getSize(50),
                  ),
                ],
              ),
              const Gap(24),
              Column(
                children: [
                  ReadOnlyField(
                    labelText: "Full Name",
                    value: isLoggedIn ? customer.name : 'Guest',
                  ),
                  const Gap(12),
                  ReadOnlyField(
                    labelText: "Date of birth",
                    value: isLoggedIn ? customer.dateOfBirth : 'N/A',
                  ),
                  const Gap(12),
                  ReadOnlyField(
                    labelText: "Phone",
                    value: isLoggedIn ? customer.phone : 'N/A',
                  ),
                  const Gap(12),
                  ReadOnlyField(
                    labelText: "Email",
                    value: isLoggedIn ? customer.email : 'guest@example.com',
                  ),
                  const Gap(12),
                  ReadOnlyField(
                    labelText: "Address",
                    value: isLoggedIn ? customer.address : 'N/A',
                  ),
                  const Gap(12),
                  ReadOnlyField(
                    labelText: "Gender",
                    value: isLoggedIn ? customer.gender : 'N/A',
                  ),
                ],
              ),
              const Gap(36),
            ],
          ),
        ),
      );
    });
  }
}
