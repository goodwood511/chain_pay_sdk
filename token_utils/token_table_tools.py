import csv
import json

def csv_to_json(csv_path, json_path):
    # 读取CSV文件并转换为字典列表
    with open(csv_path, 'r', encoding='utf-8') as csv_file:
        csv_reader = csv.DictReader(csv_file)

        # 转换数字字段类型
        data = []
        for row in csv_reader:
            # 处理数值型字段
            converted = {}
            for key, value in row.items():
                if key in ['chain_id', 'token_id', 'decimals']:
                    converted[key] = int(value) if value.strip() else None
                else:
                    converted[key] = value
            data.append(converted)

    # 写入JSON文件
    with open(json_path, 'w', encoding='utf-8') as json_file:
        json.dump(data, json_file, indent=2, ensure_ascii=False)

if __name__ == "__main__":
    csv_to_json('token_table.csv', 'output.json')
    print("转换完成！结果已保存到 output.json")